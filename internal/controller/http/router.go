package http

import (
	"github.com/flaiers/flallet/internal/config"
	v1 "github.com/flaiers/flallet/internal/controller/http/v1"
	"github.com/gofiber/fiber/v2"
)

func Router(cfg config.Config) func(fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/v1", v1.Router(cfg))
		router.Get("/healthcheck", func(ctx *fiber.Ctx) error {
			return ctx.SendStatus(fiber.StatusOK)
		})
	}
}
