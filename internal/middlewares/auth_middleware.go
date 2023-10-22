package middlewares

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

// Authourize admin
func AuthorizeAdmin(c *fiber.Ctx) error {
	fmt.Println("MW: Authorize Admin")

	tokenString := c.Cookies("Authorize Admin")

	// Check if it is admin
	isValid, claims := helpers.IsValidToken(tokenString, c)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed!",
			"message": " Unauthorized access. Invalid Token",
		})
	}
	admin := claims.(*helpers.CustomClaims).Model

	c.Locals("AdminModel", admin)

	return c.Next()
}

// Authorize restaurant
func AuthorizeRestaurant(c *fiber.Ctx) error {
	fmt.Println("MW: Authorize Restaurant")

	tokenString := c.Cookies("Authorize Restaurant")

	// Check if it is restaurant
	isValid, claims := helpers.IsValidToken(tokenString, c)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed!",
			"message": " Unauthorized access. Invalid Token",
		})
	}
	restaurant := claims.(*helpers.CustomClaims).Model

	c.Locals("RestaurantModel", restaurant)

	return c.Next()
}

// Authorize user
func AuthorizeUser(c *fiber.Ctx) error {
	fmt.Println("MW: Authorize User")

	tokenString := c.Cookies("Authorize User")

	// Check if it is user
	isValid, claims := helpers.IsValidToken(tokenString, c)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed!",
			"message": " Unauthorized access. Invalid Token",
		})
	}
	user := claims.(*helpers.CustomClaims).Model

	c.Locals("UserModel", user)

	u := c.Locals("UserModel").(map[string]interface{})

	if u["status"] == "Blocked" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed! ",
			"message": "Unauthorized access. You have been blocked!",
		})
	}

	return c.Next()
}
