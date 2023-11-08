package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App, user *handlers.UserHandler) {

	f.Post("/login", user.Login)
	f.Post("/signup", user.SignUp)
	f.Post("/verifyOtp", middlewares.AuthenticateUser, user.VerifyOtp)
	f.Post("/sendOtp", middlewares.AuthenticateUser, user.SendOtp)

	u := f.Group("/", middlewares.AuthenticateUser, middlewares.AuthorizeUser)
	u.Get("/dishes", user.GetDishesPage)
	u.Get("/dishes/:id", user.GetDish)
	u.Get("/search", user.SearchDish)

}
