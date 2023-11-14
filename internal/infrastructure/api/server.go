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

func NewServerHttp(
	adminHandler *handlers.AdminHandler,
	sellerHandler *handlers.SellerHandler,
	userHandler *handlers.UserHandler,
	cartHandler *handlers.CartHandler,
	orderHandler *handlers.OrderHandler,
) *ServerHttp {
	app := fiber.New()

	//	@securityDefinitions.apikey	Bearer
	//	@in							header
	//	@name						Authorization
	//	@description				Authentication using a JSON Web Token (JWT). The token should be included in the request header named "Authorization". The format of the header is: Authorization: Bearer <token>. Replace `<token>` with the actual JWT token.
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.AdminRoutes(app, adminHandler)
	routes.SellerRoutes(app, sellerHandler, orderHandler)
	routes.UserRoutes(app, userHandler, cartHandler, orderHandler)

	return &ServerHttp{app}
}

func (sh *ServerHttp) Start() {
	sh.app.Listen("localhost:8080")
}
