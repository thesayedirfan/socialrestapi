package apigateway

import (
	"thesayedirfan/socialapi/utils/middlerware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

var App *fiber.App
var Api fiber.Router
var V1 fiber.Router

func Start() {
	App = fiber.New()
	Api = App.Group("/api")
	V1 = Api.Group("/v1")

	App.Use(helmet.New())
	App.Use(cors.New())

	App.Use(logger.New())

	PublicRoutes()

	App.Use(middlerware.TokenValid)

	PrivateRoutes()

	App.Listen(":8080")

}
