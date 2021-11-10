package server

import (
	"github.com/gofiber/fiber/v2"
	"swift-playground/handlers"
)

func registerRoutes(app *fiber.App) {
	app.Get("/buckets", handlers.ListBuckets)
}
