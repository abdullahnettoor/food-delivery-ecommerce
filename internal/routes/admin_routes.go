package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(f *fiber.App) {

	admin := f.Group("/admin")
	// admin.Use(middlewares.AuthorizeAdmin)

	admin.Get("/login", handlers.GetAdminLogin)

	admin.Post("/login", handlers.AdminLogin)
}
