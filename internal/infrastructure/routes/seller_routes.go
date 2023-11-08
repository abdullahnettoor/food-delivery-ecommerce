package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SellerRoutes(f *fiber.App, seller *handlers.SellerHandler) {

	f.Post("/seller/register", seller.SignUp)
	f.Post("/seller/login", seller.Login)

	s := f.Group("/seller", middlewares.AuthenticateSeller)

	s.Post("/addDish", seller.CreateDish)
	s.Get("/dishes", seller.GetAllDish)
	s.Get("/dishes/:id", seller.GetDish)
	s.Put("/dishes/edit/:id", seller.UpdateDish)
	s.Delete("/dishes/delete/:id", seller.DeleteDish)

}
