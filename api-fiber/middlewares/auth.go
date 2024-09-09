package middlewares

import (
	"api-fiber/libs"
	"api-fiber/models"
	"api-fiber/services"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	// Lấy cookie từ token
	cookie := c.Cookies("token")
	if cookie == "" {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": libs.AUTH_UNSUCCESS_NO_TOKEN,
		})
	}

	// Khởi tạo và cấu hình service
	service := services.AuthService{}
	service.New(&models.Credentials{})

	// Kiểm tra xem có phải trả về message
	err := service.CheckToken(cookie)
	if err != "" {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Next()
}
