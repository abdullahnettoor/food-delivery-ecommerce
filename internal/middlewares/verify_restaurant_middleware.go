package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyRestaurant(c *fiber.Ctx) error {
	r := c.Locals("RestaurantModel").(map[string]interface{})

	switch r["Status"] {
	case "Pending":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":     "failed",
			"restaurant": c.Locals("RestaurantModel"),
			"dashboard":  "You can see dashboard after admin's verification",
		})
	case "Blocked":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":     "failed",
			"restaurant": c.Locals("RestaurantModel"),
			"dashboard":  "Your account is blocked",
		})
	case "Rejected":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":     "failed",
			"restaurant": c.Locals("RestaurantModel"),
			"dashboard":  "Your account verification is rejected",
		})
	default:
		return c.Next()
	}
}
