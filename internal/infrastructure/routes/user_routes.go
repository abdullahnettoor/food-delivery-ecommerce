package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App, user *handlers.UserHandler) {

	f.Post("/signup", user.SignUp)
	f.Post("/sendOtp", middlewares.AuthenticateUser, user.SendOtp)
	f.Post("/verifyOtp", middlewares.AuthenticateUser, user.VerifyOtp)
	f.Post("/login", user.Login)

	u := f.Group("/", middlewares.AuthenticateUser, middlewares.AuthorizeUser)
	u.Get("/dishes", user.GetDishesPage)
	u.Get("/dishes/:id", user.GetDish)
	u.Get("/search/dishes", user.SearchDish)

	u.Get("/sellers", user.GetSellersPage)
	u.Get("/sellers/:id", user.GetSeller)
	u.Get("/search/sellers", user.SearchSeller)

}
