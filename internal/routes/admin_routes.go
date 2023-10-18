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

	admin.Patch("/restaurants/verify/:id", handlers.VerifyRestaurant)

}
