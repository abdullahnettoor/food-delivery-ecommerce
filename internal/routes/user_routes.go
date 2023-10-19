package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App) {
	f.Post("/signup", handlers.UserSignUp)
	f.Post("/signup/verifyOtp", handlers.VerifyOtp)
	f.Post("/login", handlers.UserLogin)
}
