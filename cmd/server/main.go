package main

import (
	"os"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	routes.AdminRoutes(app)

	app.Listen("localhost:" + os.Getenv("PORT"))

}
