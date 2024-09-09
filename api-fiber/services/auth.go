package services

import (
	"api-fiber/config"
	"api-fiber/libs"
	"api-fiber/models"

	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	creds          models.Credentials
	jwtKey         []byte
	ExpirationTime time.Time
}

func (service *AuthService) New(creds *models.Credentials) {
	service.creds = *creds
	service.jwtKey = []byte(config.DEFAULT_SECRET_KEY)

	// Lấy dữ liệu từ TOKEN_LOGIN_TIME, nếu không thành công tạo giá trị mặc định
	expirationTime, err := strconv.Atoi(config.TOKEN_LOGIN_MIN)
	if err != nil {
		expirationTime = 30
	}
	service.ExpirationTime = time.Now().Add(time.Duration(expirationTime) * time.Minute)
}

func (service *AuthService) GenerateToken() (string, error) {
	// Tạo JWT token
	claims := &models.Claims{
		Username: service.creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: service.ExpirationTime.Unix(),
		},
	}

	// Tạo token string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(service.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) CheckToken(tokenString string) string {
	// Parse và xác thực token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra phương thức ký và trả về jwtKey
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return service.jwtKey, nil
	})
	if err != nil || !token.Valid {
		return libs.AUTH_UNSUCCESS_INVALID_TOKEN
	}

	// Kiểm tra claims của token
	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return libs.AUTH_UNSUCCESS_NO_CLAIMS
	}

	// Kiểm tra expiration time của token
	if claims.ExpiresAt < time.Now().Unix() {
		return libs.AUTH_UNSUCCESS_TOKEN_EXPIRED
	}

	// Kiểm tra username của token
	if claims.Username != config.DEFAULT_USERNAME {
		return libs.AUTH_UNSUCCESS_INFORMATION_INVALID
	}

	return ""
}
