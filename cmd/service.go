package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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

	setupRoutes(app, cfg)

	return app.Listen(":" + cfg.Port)
}
