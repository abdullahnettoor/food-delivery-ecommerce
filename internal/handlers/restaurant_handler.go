package handlers

import (
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RestuarantSignUp(c *fiber.Ctx) error {
	Body := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}{}

	c.BodyParser(&Body)

	if Body.Email == "" || Body.Password == "" || Body.Name == "" || Body.Description == "" {
		fmt.Println("All fields should be filled")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "All fields should be filled"})
	}
	fmt.Println("Finding email", Body.Email)
	result := initializers.DB.Exec(`SELECT email FROM restaurants WHERE email = ?`, Body.Email)
	if result.Error != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}
	if result.RowsAffected != 0 {
		fmt.Println("Restaurant with provided email already exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Restaurant with email entered already exist"})
	}

	hashedPassword, err := helpers.HashPassword(Body.Password)
	if err != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "Bcrypt Error", "error": err})
	}

	restaurant := models.Restaurant{
		Email:       Body.Email,
		Password:    hashedPassword,
		Name:        Body.Name,
		Description: Body.Description,
	}

	result = initializers.DB.Create(&restaurant)
	if result.Error != nil {
		fmt.Println("Error occured while creating new Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	result = initializers.DB.Raw(`SELECT * FROM restaurants WHERE email = ?`, Body.Email).Scan(&restaurant)
	if result.Error != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	token, err := helpers.CreateToken(c, "Restaurant", time.Hour*24, restaurant)
	if err != nil {
		fmt.Println("Error Creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}
	fmt.Println("Token created")
	c.Cookie(&fiber.Cookie{Name: "Authorize Restaurant", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":     "success",
		"message":    "You can add items to sell, only after the verification made by admin",
		"restaurant": c.Locals("RestaurantModel"),
	})
}

func RestaurantLogin(c *fiber.Ctx) error {
	Body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	RestaurantDetails := models.Restaurant{}
	c.BodyParser(&Body)

	fmt.Println("From Request", Body)

	if Body.Email == "" {
		fmt.Println("Email shouldn't be empty")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Email field shouldn't be empty"})
	}

	result := initializers.DB.Raw(`SELECT * FROM restaurants WHERE email = ?`, Body.Email).Scan(&RestaurantDetails)
	if result.Error != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	if result.RowsAffected < 1 {
		fmt.Println("Restaurant with provided email don't exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "No Restaurant exist with email entered"})
	}

	fmt.Println("From DB", RestaurantDetails)

	if ok, err := helpers.CompareHashedPassword(RestaurantDetails.Password, Body.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Bcrypt Error", "error": err})
	} else if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Password is wrong"})
	}

	token, err := helpers.CreateToken(c, "Restaurant", time.Hour*24, RestaurantDetails)
	if err != nil {
		fmt.Println("Error Creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}
	fmt.Println("Token created")
	c.Cookie(&fiber.Cookie{Name: "Authorize Restaurant", Value: token})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"message":    "Restaurant Logged in successfully",
		"token":      token,
		"restaurant": c.Locals("RestaurantModel"),
	})
}

func RestaurantDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":     "success",
		"restaurant": c.Locals("RestaurantModel"),
		"dashboard":  "Dasboard data will be passed here",
	})
}

func AddDish(c *fiber.Ctx) error {
	restaurant := c.Locals("RestaurantModel").(map[string]interface{})

	dish := models.Dish{}
	c.BodyParser(&dish)
	dish.RestaurantID, _ = uuid.Parse(restaurant["restaurantId"].(string))

	if 0 > dish.Price || dish.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Given datas are invalid"})
	}

	// Add new dish to DB
	dishId := initializers.DB.Create(&dish)
	if dishId.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": dishId.Error})
	}

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE id = ?`, dish.ID).Scan(&dish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": dishId.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Dish Added", "dish": dish, "restaurant": c.Locals("RestaurantModel")})
}

func GetDish(c *fiber.Ctx) error {
	restaurant := c.Locals("RestaurantModel").(map[string]interface{})
	dish := models.Dish{}
	dishId := c.Params("id")

	fmt.Println("Restaurant ID is", restaurant["restaurantId"])

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE restaurant_id = ? AND id = ? AND deleted_at IS NULL`, restaurant["restaurantId"], dishId).Scan(&dish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "dish": dish, "restaurant": c.Locals("RestaurantModel")})
}

func GetAllDishes(c *fiber.Ctx) error {
	restaurant := c.Locals("RestaurantModel").(map[string]interface{})
	dishList := []models.Dish{}

	fmt.Println("Restaurant ID is", restaurant["restaurantId"])

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE restaurant_id = ? AND deleted_at IS NULL`, restaurant["restaurantId"]).Scan(&dishList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "dishList": dishList, "restaurant": c.Locals("RestaurantModel")})
}

func EditDish(c *fiber.Ctx) error {
	restaurant := c.Locals("RestaurantModel").(map[string]interface{})
	dishId := c.Params("id")
	dbDish := models.Dish{}

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE id = ?`, dishId).Scan(&dbDish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	dish := models.Dish{}
	c.BodyParser(&dish)
	dish.ID, _ = uuid.Parse(dishId)
	dish.RestaurantID, _ = uuid.Parse(restaurant["restaurantId"].(string))

	if 0 > dish.Price || dish.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Given datas are invalid"})
	}

	// Save edits to DB
	result = initializers.DB.Save(&dish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Dish Edited Successfully", "dish": dish, "restaurant": c.Locals("RestaurantModel")})
}

func DeleteDish(c *fiber.Ctx) error {
	dishId := c.Params("id")

	fmt.Println("Dish ID to Delete", dishId)

	result := initializers.DB.Exec(`UPDATE dishes SET deleted_at = NOW() WHERE id = ?`, dishId)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed! DB Error",
			"error":  result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Dish deleted"})
}
