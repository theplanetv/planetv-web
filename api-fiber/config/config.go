package config

import (
	"context"
	"fmt"
	"os"
)

var CTX context.Context = context.Background()

var (
	// Authentication config
	DEFAULT_USERNAME   string = os.Getenv("DEFAULT_USERNAME")
	DEFAULT_PASSWORD   string = os.Getenv("DEFAULT_PASSWORD")
	DEFAULT_SECRET_KEY string = os.Getenv("DEFAULT_SECRET_KEY")
	BCRYPT_COST        string = os.Getenv("BCRYPT_COST")
	TOKEN_LOGIN_MIN    string = os.Getenv("TOKEN_LOGIN_MIN")

	// API config
	API_PORT string = os.Getenv("API_PORT")

	// Postgresql database config
	POSTGRES_USERNAME string = os.Getenv("POSTGRES_USERNAME")
	POSTGRES_PASSWORD string = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST     string = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT     string = os.Getenv("POSTGRES_PORT")
	POSTGRES_DATABASE string = os.Getenv("POSTGRES_DATABASE")
	URL_DATABASE      string = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		POSTGRES_HOST, POSTGRES_USERNAME, POSTGRES_PASSWORD, POSTGRES_DATABASE, POSTGRES_PORT,
	)
)
