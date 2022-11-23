package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/flaiers/flallet/internal/config"
	"github.com/flaiers/flallet/internal/controller/http"
)

func New(cfg config.Config) *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New(config.NewCorsConfig(cfg)))
	app.Use(logger.New(config.NewLoggerConfig()))
	app.Route(config.API, http.Router(cfg))

	return app
}
