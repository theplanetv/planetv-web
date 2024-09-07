package config

import (
	"context"
)

var CTX context.Context = context.Background()

var (
	// Postgresql database config
	DATABASE_USERNAME string = "postgres"
	DATABASE_PASSWORD string = "postgres"
	DATABASE_HOST     string = "localhost"
	DATABASE_PORT     string = "5432"
	DATABASE_DATABASE string = "planetv"
	URL_DATABASE      string = "postgres://" + DATABASE_USERNAME + ":" +
		DATABASE_PASSWORD + "@" + DATABASE_HOST + ":" +
		DATABASE_PORT + "/" + DATABASE_DATABASE + "?sslmode=disable"
)
