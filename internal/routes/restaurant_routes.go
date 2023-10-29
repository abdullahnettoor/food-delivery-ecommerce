package routes

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/handlers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RestaurantRoutes(f *fiber.App) {

	f.Post("/restaurant/register", handlers.RestuarantSignUp)
	f.Post("/restaurant/login", handlers.RestaurantLogin)

	restaurant := f.Group("/restaurant", middlewares.AuthorizeRestaurant, middlewares.VerifyRestaurant)

	restaurant.Get("/dashboard", handlers.RestaurantDashboard)
	restaurant.Post("/addDish", handlers.AddDish)
	restaurant.Get("/dishes", handlers.GetAllDishes)
	restaurant.Get("/dishes/dish/:id", handlers.GetDish)
	restaurant.Put("/dishes/edit/dish/:id", handlers.EditDish)
	restaurant.Delete("/dishes/delete/dish/:id", handlers.DeleteDish)

	restaurant.Get("/orders", handlers.ViewRestaurantOrders)
	restaurant.Get("/orders/:id", handlers.ViewRestaurantOrderDetails)
	restaurant.Patch("/orders/:id", handlers.UpdateOrderStatus)

}
