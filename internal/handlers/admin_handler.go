package handlers

import (
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func AdminLogin(c *fiber.Ctx) error {
	Body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	AdminDetails := models.Admin{}

	c.BodyParser(&Body)

	fmt.Println("From Request", Body)

	if Body.Email == "" {
		fmt.Println("Email shouldn't be empty")
		return c.JSON(fiber.Map{"status": "failed", "message": "Email field shouldn't be empty"})
	}

	result := initializers.DB.Raw(`SELECT * FROM admins WHERE email = ?`, Body.Email).Scan(&AdminDetails)

	if result.Error != nil {
		fmt.Println("Error Occured while fetching Admin", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	if result.RowsAffected < 1 {
		fmt.Println("Admin with provided email don't exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "No admin exist with email entered"})
	}

	fmt.Println("From DB", AdminDetails)

	if Body.Email != AdminDetails.Email || Body.Password != AdminDetails.Password {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Invalid Email or Password"})
	}

	token, err := helpers.CreateToken(c, "Admin", time.Hour*24, AdminDetails)
	if err != nil {
		fmt.Println("Error Creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}
	fmt.Println("Token created")
	c.Cookie(&fiber.Cookie{Name: "Authorize Admin", Value: token})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Admin Logged in successfully",
		"token":   token,
		"admin":   c.Locals("AdminModel"),
	})
}

func AdminDashboard(c *fiber.Ctx) error {
	restaurantList := []models.Restaurant{}
	result := initializers.DB.Raw(`SELECT * FROM restaurants LIMIT 5`).Scan(&restaurantList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": result.Error})
	}

	userList := []models.User{}
	res := initializers.DB.Raw(`SELECT * FROM users`).Scan(&userList)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": res.Error})
	}

	return c.JSON(fiber.Map{
		"status":         "success",
		"dashboard":      "dashboard data will be generated here",
		"admin":          c.Locals("AdminModel"),
		"restaurantList": restaurantList,
		"userList":       userList,
	})
}

func GetRestaurants(c *fiber.Ctx) error {
	query := c.Query("status")

	restaurantList := []models.Restaurant{}
	var result *gorm.DB

	if query == "Pending" || query == "Verified" || query == "Blocked" || query != "Rejected" {
		result = initializers.DB.Raw(`SELECT * FROM restaurants WHERE status ILIKE ?`, query).Scan(&restaurantList)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}
	}

	if query == "" {
		result = initializers.DB.Raw(`SELECT * FROM restaurants`).Scan(&restaurantList)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": result.Error})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":          "success",
		"message":         "Request Successful",
		"admin":           c.Locals("AdminModel"),
		"restaurantList":  restaurantList,
		"restaurantCount": result.RowsAffected,
	})
}

func GetRestaurant(c *fiber.Ctx) error {
	resId := c.Params("id")

	restaurant := models.Restaurant{}
	result := initializers.DB.Raw(`SELECT * FROM restaurants WHERE id = ?`, resId).Scan(&restaurant)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": result.Error})
	}
	if result.RowsAffected < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Restaurant Not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "admin": c.Locals("AdminModel"), "restaurant": restaurant})
}

func VerifyRestaurant(c *fiber.Ctx) error {
	resId := c.Params("id")
	fmt.Println("ID is", resId)

	result := initializers.DB.Exec(`UPDATE restaurants SET status = 'Verified' WHERE id = ?`, resId)
	if result.Error != nil {
		fmt.Println("Restaurant Verified")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Restaurant Verified Successfully"})
}

func BlockRestaurant(c *fiber.Ctx) error {
	resId := c.Params("id")

	result := initializers.DB.Exec(`UPDATE restaurants SET status = 'Blocked' WHERE id = ?`, resId)
	if result.Error != nil {
		fmt.Println("Restaurant Blocked")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Restaurant Blocked"})
}

func GetAllUsers(c *fiber.Ctx) error {
	userList := []models.User{}
	result := initializers.DB.Raw(`SELECT * FROM users`).Scan(&userList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": result.Error})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "admin": c.Locals("AdminModel"), "userCount": result.RowsAffected, "userList": userList})
}

func GetUser(c *fiber.Ctx) error {
	resId := c.Params("id")

	user := models.User{}
	result := initializers.DB.Raw(`SELECT * FROM users WHERE id = ?`, resId).Scan(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "DB Error", "error": result.Error})
	}
	if result.RowsAffected < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "User Not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "admin": c.Locals("AdminModel"), "user": user})
}

func BlockUser(c *fiber.Ctx) error {
	resId := c.Params("id")

	result := initializers.DB.Exec(`UPDATE users SET status = 'Blocked' WHERE id = ?`, resId)
	if result.Error != nil {
		fmt.Println("User Blocked")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User Blocked"})
}
