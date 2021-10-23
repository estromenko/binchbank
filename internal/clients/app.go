package clients

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *fiber.App {
	db.AutoMigrate(
		&Client{},
		&OperationType{},
		&Operation{},
	)

	controller := clientsController{db: db}

	app := fiber.New()
	{
		app.Get("/", controller.All())
		app.Get("/:id/operations", controller.Operations())
	}

	return app
}
