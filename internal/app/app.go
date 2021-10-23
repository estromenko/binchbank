package app

import (
	"os"

	"github.com/estromenko/binchbank/internal/controllers"
	"github.com/estromenko/binchbank/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
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
	app.Use(cors.New())

	app.Use(logger.New())

	app.Mount("/", controllers.NewController(db))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("JWTSECRET")),
		TokenLookup: "cookie:token",
	}))

	return app.Listen(":8888")
}
