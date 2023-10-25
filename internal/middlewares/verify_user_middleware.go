package middlewares

import "github.com/gofiber/fiber/v2"

func VerifyUser(c *fiber.Ctx) error {

	u := c.Locals("UserModel").(map[string]interface{})

	switch u["status"] {
	case "Pending":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed! ",
			"message": "Verify your account @ /verifyOtp to countinue",
		})
	case "Blocked":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed! ",
			"message": "Unauthorized access. You have been blocked!",
		})
	default:
		return c.Next()
	}

}
