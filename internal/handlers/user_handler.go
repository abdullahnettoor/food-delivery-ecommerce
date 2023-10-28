package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserSignUp(c *fiber.Ctx) error {
	user := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
	}{}

	c.BodyParser(&user)

	fmt.Println(user)

	if user.Email == "" || user.Password == "" || user.FirstName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "The fields shouldn't be empty"})
	}

	res := initializers.DB.Exec(`SELECT email FROM users WHERE email = ?`, user.Email)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": res.Error})
	}
	if res.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "user with email already exist"})
	}

	err := helpers.SendOtp(user.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "OTP Error", "error": err})
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error Occured while fetching Restaurant", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "Bcrypt Error", "error": err})
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  hashedPassword,
	}
	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}
	result.Row().Scan(&newUser)

	token, err := helpers.CreateToken(c, "User", time.Hour*24, newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}

	c.Cookie(&fiber.Cookie{Name: "Authorize User", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "verify otp @ /verifyOtp", "user": c.Locals("UserModel"), "token": token})
}

func VerifyOtp(c *fiber.Ctx) error {
	body := struct {
		OTP string `json:"otp"`
	}{}
	c.BodyParser(&body)

	u := c.Locals("UserModel").(map[string]interface{})

	status, err := helpers.VerifyOtp(fmt.Sprintf("%v", u["phone"]), body.OTP)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "OTP Error", "error": err})
	}
	if status != "approved" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "OTP is invalid"})
	}

	result := initializers.DB.Exec(`UPDATE users SET status = 'Active' WHERE id = ?`, u["userId"])
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	var user models.User
	result = initializers.DB.Raw(`SELECT * FROM users WHERE id = ?`, u["userId"]).Scan(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	token, err := helpers.CreateToken(c, "User", time.Hour*24, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}
	c.Locals("UserModel", user)
	c.Cookie(&fiber.Cookie{Name: "Authorize User", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "user": c.Locals("UserModel"), "message": "User verified successfully"})
}

func UserLogin(c *fiber.Ctx) error {
	user := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	c.BodyParser(&user)

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Fields shouldn't be empty"})
	}

	dbUser := models.User{}
	result := initializers.DB.Raw(`SELECT * FROM users WHERE email = ?`, user.Email).Scan(&dbUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "No user registered with this email"})
	}

	if ok, err := helpers.CompareHashedPassword(dbUser.Password, user.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Bcrypt Error", "error": err})
	} else if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Password is wrong"})
	}

	token, err := helpers.CreateToken(c, "User", time.Hour*24, dbUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}

	c.Cookie(&fiber.Cookie{Name: "Authorize User", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User Logged In successfully", "user": c.Locals("UserModel"), "token": token})
}

func GetDishes(c *fiber.Ctx) error {
	dishList := []models.Dish{}
	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Error occured while parsing URL params", "error": err.Error})
	}
	limit := 5
	offset := (page - 1) * int64(limit)

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE deleted_at IS NULL LIMIT ? OFFSET ?`, limit, offset).Scan(&dishList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "dishList": dishList, "user": c.Locals("UserModel")})
}

func AddToCart(c *fiber.Ctx) error {
	dishId := c.Params("id")
	dish := models.Dish{}
	dbCartItem := models.CartItem{}
	user := c.Locals("UserModel").(map[string]any)

	result := initializers.DB.Raw(`
	SELECT 
	*
	FROM dishes
	WHERE id = ? AND
	deleted_at IS NULL AND
	quantity > 0 AND
	availability = true`, dishId).Scan(&dish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Selected dish is not available",
		})
	}

	userId, err := uuid.Parse(user["userId"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "UUID Error", "error": err})
	}

	result = initializers.DB.Raw(`SELECT * FROM cart_items WHERE id = ? AND dish_id = ?`, userId, dishId).Scan(&dbCartItem)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	if result.RowsAffected != 0 {
		if dish.RestaurantID != dbCartItem.RestaurantID {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "You can't order from multiple restaurants at the same time."})
		}
		dbCartItem.Quantity = dbCartItem.Quantity + 1
		initializers.DB.Save(&dbCartItem)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":   "success",
			"message":  "Dish added to cart",
			"cartItem": dbCartItem,
			"user":     c.Locals("UserModel"),
		})
	}

	cartItem := models.CartItem{
		ID:           userId,
		RestaurantID: dish.RestaurantID,
		DishID:       dish.ID,
		Quantity:     1,
	}

	result = initializers.DB.Create(&cartItem)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "success",
		"message":  "Dish added to cart",
		"cartItem": cartItem,
		"user":     c.Locals("UserModel"),
	})
}

func ViewCart(c *fiber.Ctx) error {
	cartDishes := []struct {
		models.Dish
		Quantity uint `json:"dishQuantity" gorm:"notNull"`
	}{}
	user := c.Locals("UserModel").(map[string]any)
	userId, err := uuid.Parse(user["userId"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "UUID Error", "error": err})
	}

	result := initializers.DB.Raw(`
	SELECT 
	d.*,
	c.quantity
	FROM cart_items c 
	INNER JOIN dishes d 
	ON d.id = c.dish_Id 
	WHERE c.id = ?
	`, userId).Scan(&cartDishes)

	if result.Error != nil {
		fmt.Println(result.Error)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "cart": "Cart is Empty"})
	}

	var totalPrice float64
	for _, dish := range cartDishes {
		totalPrice += (float64(dish.Price) * float64(dish.Quantity))
	}

	cart := models.Cart{
		ID:           userId,
		RestaurantID: cartDishes[0].RestaurantID,
		Dishes:       cartDishes,
		TotalPrice:   totalPrice,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "cart": cart, "user": c.Locals("UserModel")})
}

func DeleteCartItem(c *fiber.Ctx) error {
	dishId := c.Params("id")
	user := c.Locals("UserModel").(map[string]any)

	result := initializers.DB.Exec(`DELETE FROM cart_items WHERE id = ? and dish_id = ?`, user["userId"], dishId)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Dish deleted from cart successfully.", "user": user})
}
