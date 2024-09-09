package services

import (
	"api-fiber/libs"
	"api-fiber/models"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateToken_AuthService(t *testing.T) {
	// Khai báo dữ liệu
	service := AuthService{}

	t.Run("Generate token unsuccess", func(t *testing.T) {
		// Lấy token
		_, err := service.GenerateToken()

		// Kiểm tra lỗi
		assert.NoError(t, err, libs.TEST_HAVE_ERROR)
	})

	t.Run("Generate token success", func(t *testing.T) {
		// Khởi tạo dữ liệu
		creds := models.Credentials{Username: "admin", Password: "admin_password"}
		service.New(&creds)

		// Lấy token
		token, err := service.GenerateToken()

		// Kiểm tra lỗi
		assert.NoError(t, err, libs.TEST_HAVE_NO_ERROR)
		assert.NotEmpty(t, token, libs.TEST_DATA_NOT_EMPTY)
	})
}
