package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/odesiuk/pg-stats-api/internal/query"
	"github.com/odesiuk/pg-stats-api/internal/storage/repositories"
	"github.com/odesiuk/pg-stats-api/pkg/db"
	"github.com/odesiuk/pg-stats-api/pkg/rest"
)

// Start setup API service.
func Start(cfg Config) error {
	app := fiber.New(fiber.Config{
		// add custom errors.
		ErrorHandler: rest.ErrorHandler,
	})

	// request ID middleware.
	app.Use(requestid.New())

	// request logger middleware.
	app.Use(logger.New())

	if err := setup(app, cfg); err != nil {
		return err
	}

	return app.Listen(":" + cfg.Port)
}

// healthcheck returns service name.
func healthcheck(cfg Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(rest.H{
			"name": cfg.AppName,
		})
	}
}

// setup init.
func setup(app *fiber.App, cfg Config) error {
	// get DB connection.
	dbConn, err := db.NewConnectionFromENV("PG")
	if err != nil {
		return err
	}

	// init controllers.
	qc := query.NewController(
		repositories.NewPgStatStatementRepo(dbConn),
		cfg.MinQueryDuration,
	)

	// add routes.
	app.Get("/healthcheck", healthcheck(cfg))
	app.Get("/queries", qc.GetAll)

	return nil
}
