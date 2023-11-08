package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(f *fiber.App, admin *handlers.AdminHandler) {

	f.Post("/admin/login", admin.Login)

	r := f.Group("/admin", middlewares.AuthenticateAdmin)

	r.Get("/sellers", admin.GetAllSellers)
	r.Patch("/sellers/verify/:id", admin.VerifySeller)
	r.Patch("/sellers/block/:id", admin.BlockSeller)
	r.Patch("/sellers/unblock/:id", admin.UnblockSeller)

	r.Get("/users", admin.GetAllUsers)
	r.Patch("/users/block/:id", admin.BlockUser)
	r.Patch("/users/unblock/:id", admin.UnblockUser)

	r.Post("/category/addCategory", admin.AddCategory)
	r.Patch("/category/edit/:id", admin.EditCategory)

}
