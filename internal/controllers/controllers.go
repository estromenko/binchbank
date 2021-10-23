package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type clientController struct {
	db *gorm.DB
}

func (c *clientController) route() *fiber.App {
	app := fiber.New()

	app.Get("/", c.Info())
	app.Get("/circle", c.CircleDiagramInfo())

	return app
}

func (c *clientController) Info() fiber.Handler {
	type response struct {
		Username        string    `json:"username"`
		CreatedAt       time.Time `json:"created_at"`
		OperationsCount uint      `json:"operations_count"`
		OperationsPlus  int       `json:"operations_plus"`
	}

	return func(ctx *fiber.Ctx) error {
		rows, err := c.db.Table("clients").Select(
			"clients.username, clients.created_at, count(*) as operations_count, sum(operations.amount) as operations_plus",
		).Joins("join operations on clients.id = operations.client_id").Group("clients.username, clients.created_at").Rows()

		if err != nil {
			return ctx.Status(400).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		var res []*response

		for rows.Next() {
			var rowData response
			if err := rows.Scan(&rowData); err != nil {
				return ctx.Status(400).JSON(map[string]string{
					"error": err.Error(),
				})
			}

			res = append(res, &rowData)
		}

		return ctx.Status(200).JSON(res)
	}
}

func (c *clientController) CircleDiagramInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var res map[string]interface{}
		c.db.Table("operations").Select("count(operations.type)").Group("operations.type").Scan(&res)

		return ctx.Status(200).JSON(res)
	}
}

func NewClientController(db *gorm.DB) *fiber.App {
	controller := clientController{db: db}
	return controller.route()
}
