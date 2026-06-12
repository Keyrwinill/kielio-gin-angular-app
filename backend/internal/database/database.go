package database

import (
	"context"
	"fmt"

	"kielio-gin-angular-app/backend/internal/config"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_SSLMODE"),
	)

	return pgx.Connect(context.Background(), connString)
}