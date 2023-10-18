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
			"failed": "Unauthorized access",
		})
	}
	admin := claims.(*helpers.CustomClaims).Model

	c.Locals("userModel", admin)

	return c.Next()
}

// // Authorize restaurant
// func AuthorizeRestaurant(c *fiber.Ctx) error {
// 	fmt.Println("MW: Authorize Restaurant")

// 	tokenString := c.Cookies("Authorize Restaurant")

// 	// Check if it is restaurant
// 	isValid := helpers.IsValidToken(tokenString, c)
// 	if !isValid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"failed": "Unauthorized access",
// 		})
// 	}

// 	return c.Next()
// }

// // Authorize user
// func AuthorizeUser(c *fiber.Ctx) error {
// 	fmt.Println("MW: Authorize User")

// 	tokenString := c.Cookies("Authorize User")

// 	// Check if it is user
// 	isValid := helpers.IsValidToken(tokenString, c)
// 	if !isValid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"failed": "Unauthorized access",
// 		})
// 	}

// 	return c.Next()
// }

