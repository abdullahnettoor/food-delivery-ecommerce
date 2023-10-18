package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RestaurantRoutes(f *fiber.App) {

	f.Post("/restaurant/register", handlers.RestuarantSignUp)
	f.Post("/restaurant/login", handlers.RestaurantLogin)

	restaurant := f.Group("/restaurant", middlewares.AuthorizeRestaurant)

	restaurant.Get("/dashboard", handlers.RestaurantDashboard)
}
