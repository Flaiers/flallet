package config

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flaiers/flallet/internal/controller/http/v1/reports"
	"github.com/flaiers/flallet/internal/controller/http/v1/reservations"
	"github.com/flaiers/flallet/internal/controller/http/v1/transactions"
	"github.com/flaiers/flallet/internal/controller/http/v1/users"
	"github.com/flaiers/flallet/internal/service"
	"github.com/flaiers/flallet/internal/usecase/utils"
	"gorm.io/gorm"

	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/storage/s3"
	"github.com/gofiber/swagger"
	"github.com/swaggo/swag"
)

const (
	API = "/api"
)

type Config struct {
	Port string `env:"BACKEND_PORT"`
	Host string `env:"BACKEND_HOST"`
	Addr string

	CorsOrigins string `env:"BACKEND_CORS_ORIGINS"`

	Database struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Pass     string `env:"DB_PASS"`
		Name     string `env:"DB_NAME"`
		TimeZone string `env:"DB_TIMEZONE" envDefault:"UTC"`
		SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
		Dsn      string
	}

	S3 struct {
		Region          string        `env:"S3_REGION"`
		Bucket          string        `env:"S3_BUCKET"`
		Endpoint        string        `env:"S3_ENDPOINT"`
		AccessKey       string        `env:"S3_ACCESS_KEY"`
		SecretAccessKey string        `env:"S3_SECRET_ACCESS_KEY"`
		MaxAttempts     int           `env:"S3_MAX_ATTEMPTS" envDefault:"3"`
		RequestTimeout  time.Duration `env:"S3_REQUEST_TIMEOUT" envDefault:"0"`
	}
}

func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	if cfg.Addr == "" {
		cfg.Addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	}

	if cfg.Database.Dsn == "" {
		cfg.Database.Dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Pass,
			cfg.Database.Name,
			cfg.Database.SSLMode,
			cfg.Database.TimeZone,
		)
	}

	return cfg, nil
}

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  utils.JSONErrorHandler,
	}
}

func NewCorsConfig(cfg Config) cors.Config {
	return cors.Config{
		AllowOrigins:     cfg.CorsOrigins,
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}
}

func NewLoggerConfig() logger.Config {
	return logger.ConfigDefault
}

func NewSwaggerConfig(s *swag.Spec) swagger.Config {
	return swagger.Config{Title: s.Title}
}

func NewS3Config(cfg Config) s3.Config {
	return s3.Config{
		Region:         cfg.S3.Region,
		Bucket:         cfg.S3.Bucket,
		Endpoint:       cfg.S3.Endpoint,
		MaxAttempts:    cfg.S3.MaxAttempts,
		RequestTimeout: cfg.S3.RequestTimeout,
		Credentials: s3.Credentials{
			AccessKey:       cfg.S3.AccessKey,
			SecretAccessKey: cfg.S3.SecretAccessKey,
		},
	}
}

func NewUsersConfig(db *gorm.DB) users.Config {
	return users.Config{
		Service: service.NewUserService(db),
	}
}

func NewReportsConfig(db *gorm.DB, storage *s3.Storage) reports.Config {
	return reports.Config{
		Service: service.NewReportService(db, storage),
	}
}

func NewReservationsConfig(db *gorm.DB) reservations.Config {
	return reservations.Config{
		Service: service.NewReservationService(db),
	}
}

func NewTransactionsConfig(db *gorm.DB) transactions.Config {
	return transactions.Config{
		Service: service.NewTransactionService(db),
	}
}
