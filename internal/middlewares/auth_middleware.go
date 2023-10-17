package middlewares

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

// Authourize admin
func AuthorizeAdmin(c *fiber.Ctx) error {
	fmt.Println("In Admin MiddleWare")

	tokenString := c.Cookies("Authorize Admin")

	// Check if it is admin
	isValid := helpers.IsValidToken(tokenString, c)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"failed": "Unauthorized access",
		})
	}

	return c.Next()
}

// Authorize restaurant
func AuthorizeRestaurant(c *fiber.Ctx) {

}

// Authorize user
func AuthorizeUser(c *fiber.Ctx) {

}
