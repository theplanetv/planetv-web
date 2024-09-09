package controllers

import (
	"api-fiber/config"
	"api-fiber/libs"
	"io"
	"time"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test the Login function without mocks
func TestLogin(t *testing.T) {
	// Setup fiber and route login
	app := fiber.New(fiber.Config{
		IdleTimeout: 15 * time.Second, // Set a global timeout of 5 seconds
	})
	app.Post("/login", Login)

	t.Run("Login unsuccess invalid information", func(t *testing.T) {
		// Mock request data
		body := `{"username":"` + "test" + `","password":"` + "test" + `"}`

		// Create a new HTTP POST request
		req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Execute the request
		res, err := app.Test(req, -1)
		require.NoError(t, err)

		// Assert response status
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)

		// Read the response body
		bodyBytes, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		// Convert the body to a string
		bodyString := string(bodyBytes)

		// Check if the response contains the expected message
		assert.Contains(t, bodyString, libs.AUTH_UNSUCCESS_INFORMATION_INVALID)
	})

	t.Run("Login success", func(t *testing.T) {
		// Mock request data
		body := `{"username":"` + config.DEFAULT_USERNAME + `","password":"` + config.DEFAULT_PASSWORD + `"}`

		// Create a new HTTP POST request
		req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Execute the request
		res, err := app.Test(req, 10000)
		require.NoError(t, err)

		// Assert response status
		assert.Equal(t, http.StatusOK, res.StatusCode)

		// Read the response body
		bodyBytes, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		// Convert the body to a string
		bodyString := string(bodyBytes)

		// Check if the response contains the expected message
		assert.Contains(t, bodyString, libs.AUTH_LOGIN_SUCCESS)

		// Assert that the cookie was set
		cookies := res.Cookies()
		require.NotNil(t, cookies)
		assert.Equal(t, "auth-token", cookies[0].Name)
		require.NotEmpty(t, cookies[0].Value)
	})
}
