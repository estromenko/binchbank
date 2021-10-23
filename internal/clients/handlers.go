package clients

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type clientsController struct {
	db *gorm.DB
}

func (c *clientsController) All() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var clients []*Client
		c.db.Find(&clients)
		return ctx.JSON(clients)
	}
}

func (c *clientsController) Operations() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		var operations []*Operation

		c.db.Find(&operations, "client_id=?", id)
		return ctx.JSON(operations)
	}
}
