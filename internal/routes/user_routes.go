package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App) {
	f.Post("/signup", handlers.UserSignUp)
	f.Post("/login", handlers.UserLogin)
	f.Post("/verifyOtp", middlewares.AuthorizeUser, handlers.VerifyOtp)

	user := f.Group("/", middlewares.AuthorizeUser, middlewares.VerifyUser)
	user.Get("/dishes", handlers.GetDishes)

	user.Post("/addToCart/:id", handlers.AddToCart)
}
