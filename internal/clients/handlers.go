package clients

import (
	"context"
	"encoding/json"

	"github.com/estromenko/binchbank/models"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Index(c *fiber.Ctx) error {
	clients, err := models.Clients().All(context.Background(), boil.GetContextDB())
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return json.NewEncoder(c.Response().BodyWriter()).Encode(clients)
}
