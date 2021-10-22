package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func Run() error {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	app := fiber.New()

	return app.Listen(":8888")
}
