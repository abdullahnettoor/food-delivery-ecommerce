package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(f *fiber.App, admin *handlers.AdminHandler) {

	f.Post("/admin/login", admin.Login)

}
