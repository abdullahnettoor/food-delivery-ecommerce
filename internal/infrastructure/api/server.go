package api

import (
	_ "github.com/abdullahnettoor/food-delivery-eCommerce/docs"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type ServerHttp struct {
	app *fiber.App
}

// "securityDefinitions": {
// 	"Bearer": {
// 	"type": "apiKey",
// 	"name": "Authorization",
// 	"in": "header"
// 	}
// },

func NewServerHttp(adminHandler *handlers.AdminHandler, sellerHandler *handlers.SellerHandler, userHandler *handlers.UserHandler) *ServerHttp {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.AdminRoutes(app, adminHandler)
	routes.SellerRoutes(app, sellerHandler)
	routes.UserRoutes(app, userHandler)

	return &ServerHttp{app}
}

func (sh *ServerHttp) Start() {
	sh.app.Listen("localhost:8080")
}
