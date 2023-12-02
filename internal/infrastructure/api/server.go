package api

import (
	"net/http"

	embedfiles "github.com/abdullahnettoor/food-delivery-eCommerce"
	_ "github.com/abdullahnettoor/food-delivery-eCommerce/docs"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
)

type ServerHttp struct {
	app *fiber.App
}

func NewServerHttp(
	adminHandler *handlers.AdminHandler,
	sellerHandler *handlers.SellerHandler,
	userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	dishHandler *handlers.DishHandler,
	cartHandler *handlers.CartHandler,
	orderHandler *handlers.OrderHandler,
	offerHandler *handlers.OfferHandler,
	favHandler *handlers.FavHandler,
	couponHandler *handlers.CouponHandler,
) *ServerHttp {

	f := embedfiles.Tmpl
	views := html.New("internal/view", ".html")
	app := fiber.New(fiber.Config{Views: views})
	// app := fiber.New()

	

	app.Use(logger.New(logger.Config{TimeFormat: "2006/01/02 15:04:05"}))

	//	@title			FoodieBuddie API
	//	@description	API for managing and interacting with FoodieBuddie services.

	//	@securityDefinitions.apikey	Bearer
	//	@in							header
	//	@name						Authorization
	//	@description				Authentication using a JSON Web Token (JWT). The token should be included in the request header named "Authorization". The format of the header is: Authorization: Bearer <token>. Replace `<token>` with the actual JWT token.
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use("/internal/view", filesystem.New(filesystem.Config{Root: http.FS(f),PathPrefix: "/internal/view"}))
	// app.Use("", filesystem.New(filesystem.Config{
	// 	Root:       http.FS(env),
	// 	PathPrefix: ".",
	// 	Browse:     true,
	// }))

	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.AdminRoutes(app, adminHandler, couponHandler)
	routes.SellerRoutes(app, sellerHandler, orderHandler, dishHandler, offerHandler)
	routes.UserRoutes(app, userHandler, sellerHandler, categoryHandler, dishHandler, cartHandler, orderHandler, offerHandler, favHandler, couponHandler)

	return &ServerHttp{app}
}

func (sh *ServerHttp) Start() {
	sh.app.Listen("localhost:8080")
}
