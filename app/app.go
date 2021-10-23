package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() error {
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_DSN")), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return err
	}

	_ = db

	app := fiber.New()

	return app.Listen(":8888")
}
