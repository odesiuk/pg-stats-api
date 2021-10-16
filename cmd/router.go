package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/odesiuk/pg-stats-api/pkg/rest"
)

// index main page.
func index(cfg Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(rest.H{
			"name": cfg.AppName,
		})
	}
}

// setupRoutes add all routes.
func setupRoutes(app *fiber.App, cfg Config) {
	app.Get("/", index(cfg))
}
