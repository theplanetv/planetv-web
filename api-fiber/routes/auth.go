package routes

import (
	"api-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Get("/check-login", controllers.CheckLogin)
	auth.Delete("/logout", controllers.Logout)
}
