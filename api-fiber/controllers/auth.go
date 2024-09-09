package controllers

import (
	"api-fiber/config"
	"api-fiber/libs"
	"api-fiber/models"
	"api-fiber/services"

	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	creds := models.Credentials{}
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": libs.AUTH_UNSUCCESS_INFORMATION_INVALID,
		})
	}

	// Check username
	if creds.Username != config.DEFAULT_USERNAME || creds.Password != config.DEFAULT_PASSWORD {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": libs.AUTH_UNSUCCESS_INFORMATION_INVALID,
		})
	}

	// Get bcryptCost
	bcryptCost, err := strconv.Atoi(config.BCRYPT_COST)
	if err != nil {
		bcryptCost = bcrypt.DefaultCost
	}

	// Create hashed password from input password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcryptCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": libs.SERVER_ERROR,
		})
	}

	// Compare default password and hashed password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(config.DEFAULT_PASSWORD))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": libs.AUTH_UNSUCCESS_INFORMATION_INVALID,
		})
	}

	// Initialize Auth service
	service := services.AuthService{}
	service.New(&creds)

	// Generate token
	token, err := service.GenerateToken()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": libs.AUTH_UNSUCCESS_CANT_GENERATE_TOKEN,
		})
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "auth-token",
		Value:    token,
		Expires:  service.ExpirationTime,
		Path:     "/",
		Secure:   false,
		HTTPOnly: true,
		SameSite: "Lax",
	})

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": libs.AUTH_LOGIN_SUCCESS,
	})
}

func CheckLogin(c *fiber.Ctx) error {
	// Get auth token from cookie
	cookie := c.Cookies("auth-token")
	if cookie == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": libs.AUTH_UNSUCCESS_NO_TOKEN,
		})
	}

	// Initialize Auth service
	service := services.AuthService{}
	service.New(&models.Credentials{})

	// Check if there is returned message
	errMessage := service.CheckToken(cookie)
	if errMessage != "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": errMessage,
		})
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": libs.AUTH_SUCCESS,
	})
}

func Logout(c *fiber.Ctx) error {
	// Delete token cookie
	c.Cookie(&fiber.Cookie{
		Name:  "auth-token",
		Value: "",
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": libs.AUTH_LOGOUT_SUCCESS,
	})
}
