package clients

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	app := fiber.New()
	{
		app.Get("/index", Index)
	}

	return app
}
