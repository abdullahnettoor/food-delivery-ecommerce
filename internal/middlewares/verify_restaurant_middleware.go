package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func VerifyRestaurant(c *fiber.Ctx) error {
	r := c.Locals("RestaurantModel").(map[string]interface{})

	fmt.Println("MW: Verifiying Restaurant\nRestaurant is ", r)

	switch r["status"] {
	case "Pending":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":     "failed",
			"restaurant": c.Locals("RestaurantModel"),
			"message":    "You can add products only after admin's verification",
		})
	case "Blocked":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":     "failed",
			"restaurant": c.Locals("RestaurantModel"),
			"message":    "Your account is blocked",
		})
	case "Rejected":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":     "failed",
			"restaurant": c.Locals("RestaurantModel"),
			"message":    "Your account verification is rejected",
		})
	default:
		return c.Next()
	}
}
