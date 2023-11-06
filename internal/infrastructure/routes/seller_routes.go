package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/gofiber/fiber/v2"
)

func SellerRoutes(f *fiber.App, seller *handlers.SellerHandler) {

	f.Post("/seller/register", seller.SignUp)
	f.Post("/seller/login", seller.Login)

}
