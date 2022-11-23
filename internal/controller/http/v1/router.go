package v1

import (
	docs "github.com/flaiers/flallet/docs/v1"
	"github.com/flaiers/flallet/internal/config"
	"github.com/flaiers/flallet/internal/controller/http/v1/reports"
	"github.com/flaiers/flallet/internal/controller/http/v1/reservations"
	"github.com/flaiers/flallet/internal/controller/http/v1/transactions"
	"github.com/flaiers/flallet/internal/controller/http/v1/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/s3"
	"github.com/gofiber/swagger"
)

// @title          Flallet API
// @version        1.0
// @description    REST API for working with user balance
// @termsOfService https://swagger.io/terms

// @contact.name  Maxim Bigin
// @contact.email i@flaiers.me
// @contact.url   https://flaiers.me

// @license.name Apache 2.0
// @license.url  https://www.apache.org/licenses/LICENSE-2.0

// @basePath /api/v1

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization

func Router(cfg config.Config) func(fiber.Router) {
	db := config.NewDatabaseConnection(cfg.Database.Dsn)
	storage := s3.New(config.NewS3Config(cfg))

	return func(router fiber.Router) {
		router.Get("/swagger/*", swagger.New(
			config.NewSwaggerConfig(docs.SwaggerInfo),
		))
		router.Route("/users", users.New(
			config.NewUsersConfig(db),
		))
		router.Route("reports", reports.New(
			config.NewReportsConfig(db, storage),
		))
		router.Route("/reservations", reservations.New(
			config.NewReservationsConfig(db),
		))
		router.Route("/transactions", transactions.New(
			config.NewTransactionsConfig(db),
		))
	}
}
