package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(
	f *fiber.App,
	user *handlers.UserHandler,
	category *handlers.CategoryHandler,
	dish *handlers.DishHandler,
	cart *handlers.CartHandler,
	order *handlers.OrderHandler,
	offer *handlers.OfferHandler,
	fav *handlers.FavHandler,
	coupon *handlers.CouponHandler,
) {

	f.Post("/signup", user.SignUp)
	f.Post("/sendOtp", middlewares.AuthenticateUser, user.SendOtp)
	f.Post("/verifyOtp", middlewares.AuthenticateUser, user.VerifyOtp)
	f.Post("/login", user.Login)
	f.Post("/forgotPassword", user.ForgotPassword)
	f.Post("/resetPassword", user.ResetPassword)

	f.Get("/cart/checkout/online", order.PlaceOrderPayOnline)
	f.Post("/cart/checkout/online", order.VerifyPayment)

	f.Get("/categories", category.GetAllCategories)
	f.Get("/categories/:id", category.GetCategory)

	f.Get("/dishes", dish.GetDishesPage)
	f.Get("/dishes/:id", dish.GetDish)

	f.Get("/offers", offer.GetAllOffers)

	f.Get("user/sellers", user.GetSellersPage)
	f.Get("user/sellers/:id", user.GetSeller)

	f.Get("/search/dishes", dish.SearchDish)
	f.Get("/search/sellers", user.SearchSeller)

	u := f.Group("/", middlewares.AuthenticateUser, middlewares.AuthorizeUser)

	u.Get("/profile", user.ViewUserDetails)
	u.Patch("/profile/edit", user.EditUserDetails)
	u.Patch("/profile/changePassword", user.ChangePassword)

	u.Post("/profile/addAddress", user.AddAddress)
	u.Get("/profile/address", user.ViewAllAddress)
	u.Get("/profile/address/:id", user.ViewAddress)
	u.Put("/profile/address/:id", user.EditAddress)

	u.Post("/addToFavourite/:id", fav.AddFavItem)
	u.Get("/favourites", fav.ViewFavItems)
	u.Delete("/favourites/:id/delete", fav.DeleteFavItem)

	u.Post("/addToCart/:id", cart.AddToCart)
	u.Get("/cart", cart.ViewCart)
	u.Delete("/cart/:id/deleteItem", cart.DeleteCartItem)
	u.Patch("/cart/:id/decrement", cart.DecrementCartItem)
	u.Delete("/cart/empty", cart.EmptyCart)

	u.Get("/coupons", coupon.GetAllCouponsForUser)
	u.Get("/coupons/available", coupon.GetAvailableCouponsForUser)
	u.Get("/coupons/redeemed", coupon.GetRedeemedByUser)

	u.Post("/cart/checkout", order.PlaceOrder)
	// u.Post("/cart/checkout/online", order.PlaceOrderPayOnline)
	u.Get("/orders", order.ViewOrdersForUser)
	u.Get("/orders/:id", order.ViewOrder)
}
