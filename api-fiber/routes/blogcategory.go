package routes

import (
	"api-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func BlogcategoryRoutes(api fiber.Router) {
	controller := controllers.BlogCategoryController{}
	auth := api.Group("/blogcategory")
	auth.Get("/count", controller.Count)
	auth.Get("/", controller.GetAll)
	auth.Post("/", controller.Create)
	auth.Put("/", controller.Update)
	auth.Delete("/:id", controller.Remove)
}
