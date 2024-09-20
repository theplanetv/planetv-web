package routes

import (
	"api-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func BlogTagRoutes(api fiber.Router) {
	controller := controllers.BlogTagController{}
	auth := api.Group("/blogtag")
	auth.Get("/count", controller.Count)
	auth.Get("/", controller.GetAll)
	auth.Post("/", controller.Create)
	auth.Put("/", controller.Update)
	auth.Delete("/:id", controller.Remove)
}
