package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(f *fiber.App) {

	f.Post("/admin/login", handlers.AdminLogin)

	admin := f.Group("/admin", middlewares.AuthorizeAdmin)

	admin.Get("/dashboard", handlers.AdminDashboard)

	admin.Get("/restaurants", handlers.GetRestaurants)
	admin.Get("/restaurants/restaurant/:id", handlers.GetRestaurant)
	admin.Patch("/restaurants/verify/restaurant/:id", handlers.VerifyRestaurant)
	admin.Patch("/restaurants/block/restaurant/:id", handlers.BlockRestaurant)

	admin.Get("/users", handlers.GetAllUsers)
	admin.Get("/users/user/:id", handlers.GetUser)
	admin.Patch("/users/block/user/:id", handlers.BlockUser)
}
