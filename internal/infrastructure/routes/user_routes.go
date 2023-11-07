package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares.go"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App, user *handlers.UserHandler) {

	f.Post("/login", user.Login)
	f.Post("/signup", user.SignUp)
	f.Post("/verifyOtp", middlewares.AuthorizeUser, user.VerifyOtp)
	f.Post("/sendOtp", middlewares.AuthorizeUser, user.SendOtp)

}
