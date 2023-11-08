package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AuthorizeSeller(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]interface{})

	fmt.Println("MW: Verifiying Seller\nSeller is", seller)

	switch seller["status"] {
	case "Pending":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "You can add products only after admin's verification",
		})
	case "Blocked":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "Your account is blocked",
		})
	case "Rejected":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "Your account verification is rejected",
		})
	default:
		return c.Next()
	}
}

func AuthorizeUser(c *fiber.Ctx) error {

	u := c.Locals("UserModel").(map[string]interface{})

	switch u["status"] {
	case "Pending":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "Verify your account @ /verifyOtp to countinue",
		})
	case "Blocked":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "Unauthorized access. You have been blocked!",
		})
	default:
		return c.Next()
	}

}
