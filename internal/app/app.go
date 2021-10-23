package app

import (
	"os"

	"github.com/estromenko/binchbank/internal/controllers"
	"github.com/estromenko/binchbank/internal/models"
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

	if err := models.Migrate(db); err != nil {
		return err
	}

	app := fiber.New()

	app.Mount("/clients", controllers.NewClientController(db))

	return app.Listen(":8888")
}
