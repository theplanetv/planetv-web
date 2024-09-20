package main

import (
	"api-fiber/config"
	"api-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Start fiber
	app := fiber.New()

	// Use built-in middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	// Add routes
	api := app.Group("/api")
	routes.AuthRoutes(api)
	routes.BlogcategoryRoutes(api)

	app.Listen(":" + config.API_PORT)
}
