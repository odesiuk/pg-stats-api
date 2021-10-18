package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/odesiuk/pg-stats-api/internal/query"
	"github.com/odesiuk/pg-stats-api/internal/storage/repositories"
	"github.com/odesiuk/pg-stats-api/pkg/rest"
	"gorm.io/gorm"
)

// Setup configure API service.
func Setup(db *gorm.DB, cfg Config) *fiber.App {
	app := fiber.New(fiber.Config{
		// add custom errors.
		ErrorHandler: rest.ErrorHandler,
	})

	// request ID middleware.
	app.Use(requestid.New())

	// request logger middleware.
	app.Use(logger.New())

	// init controllers.
	qc := query.NewController(
		repositories.NewPgStatStatementRepo(db),
		cfg.MinQueryDuration,
	)

	// add routes.
	app.Get("/", index(cfg))
	app.Get("/queries", qc.GetAll)

	return app
}

// index returns service name.
func index(cfg Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(rest.H{
			"name": cfg.AppName,
		})
	}
}
