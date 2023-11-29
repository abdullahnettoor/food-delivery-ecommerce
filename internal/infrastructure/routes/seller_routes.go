package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SellerRoutes(
	f *fiber.App,
	seller *handlers.SellerHandler,
	order *handlers.OrderHandler,
	dish *handlers.DishHandler,
	offer *handlers.OfferHandler,
) {

	f.Post("/seller/register", seller.SignUp)
	f.Post("/seller/login", seller.Login)

	s := f.Group("/seller", middlewares.AuthenticateSeller, middlewares.AuthorizeSeller)

	s.Post("/addDish", dish.CreateDish)
	s.Get("/dishes", dish.GetAllDishBySeller)
	s.Get("/dishes/:id", dish.GetDishBySeller)
	s.Put("/dishes/:id", dish.UpdateDish)
	s.Delete("/dishes/:id", dish.DeleteDish)

	s.Get("/orders", order.ViewOrdersForSeller)
	s.Get("/orders/:id", order.ViewOrder)
	s.Patch("/orders/:id", order.UpdateOrderStatus)

	s.Get("/offers", offer.GetOffersBySeller)
	s.Post("/offers/addOffer", offer.CreateOffer)
	s.Put("/offers/:id", offer.UpdateOffer)
	s.Patch("/offers/:id", offer.UpdateOfferStatus)

	s.Get("/sales/daily", order.GetDailySales)
}
