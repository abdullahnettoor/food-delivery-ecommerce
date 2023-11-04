package api

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
)

type ServerHttp struct {
	app *fiber.App
}

func NewServerHttp(handler *handlers.AdminHandler) *ServerHttp {
	app := fiber.New()

	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.AdminRoutes(app, handler)

	return &ServerHttp{app}
}

func (sh *ServerHttp) Start() {
	sh.app.Listen("localhost:8080")
}
