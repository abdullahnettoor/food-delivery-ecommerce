package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App) {
	f.Post("/signup", handlers.UserSignUp)
	f.Post("/login", handlers.UserLogin)
	f.Post("/verifyOtp", middlewares.AuthorizeUser, handlers.VerifyOtp)

	user := f.Group("/", middlewares.AuthorizeUser, middlewares.VerifyUser)
	user.Get("/dishes", handlers.GetDishes)

	user.Post("/addToCart/:id", handlers.AddToCart)
	user.Get("/cart", handlers.ViewCart)
	user.Delete("/cart/delete/:id", handlers.DeleteCartItem)
	user.Patch("/cart/decrement/:id", handlers.DecrementCartItem)

	user.Post("/profile/addAddress", handlers.AddAddress)
	user.Post("/cart/checkout", handlers.PlaceOrder)
	user.Get("/profile/orders", handlers.ViewUserOrders)
	user.Get("/profile/orders/:id", handlers.ViewUserOrder)
	user.Patch("/profile/orders/cancel/:id", handlers.CancelUserOrder)

	user.Patch("/profile/changePassword", handlers.ChangeUserPassword)
}
