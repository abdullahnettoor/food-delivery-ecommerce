package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(
	f *fiber.App,
	user *handlers.UserHandler,
	dish *handlers.DishHandler,
	cart *handlers.CartHandler,
	order *handlers.OrderHandler,
	offer *handlers.OfferHandler,
) {

	f.Post("/signup", user.SignUp)
	f.Post("/sendOtp", middlewares.AuthenticateUser, user.SendOtp)
	f.Post("/verifyOtp", middlewares.AuthenticateUser, user.VerifyOtp)
	f.Post("/login", user.Login)
	f.Get("/cart/checkout/online", order.PlaceOrderPayOnline)
	f.Post("/cart/checkout/online", order.VerifyPayment)

	f.Get("/dishes", dish.GetDishesPage)
	f.Get("/dishes/:id", dish.GetDish)

	f.Get("/offers", offer.GetAllOffers)

	f.Get("user/sellers", user.GetSellersPage)
	f.Get("user/sellers/:id", user.GetSeller)

	f.Get("/search/dishes", dish.SearchDish)
	f.Get("/search/sellers", user.SearchSeller)

	u := f.Group("/", middlewares.AuthenticateUser, middlewares.AuthorizeUser)

	u.Post("/profile/addAddress", user.AddAddress)
	u.Get("/profile/address", user.ViewAllAddress)
	u.Get("/profile/address/:id", user.ViewAddress)
	u.Put("/profile/address/:id", user.EditAddress)

	u.Post("/addToCart/:id", cart.AddToCart)
	u.Get("/cart", cart.ViewCart)
	u.Delete("/cart/:id/deleteItem", cart.DeleteCartItem)
	u.Patch("/cart/:id/decrement", cart.DecrementCartItem)
	u.Delete("/cart/empty", cart.EmptyCart)

	u.Post("/cart/checkout", order.PlaceOrder)
	// u.Post("/cart/checkout/online", order.PlaceOrderPayOnline)
	u.Get("/orders", order.ViewOrdersForUser)
	u.Get("/orders/:id", order.ViewOrder)
}
