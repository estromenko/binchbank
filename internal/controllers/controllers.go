package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/estromenko/binchbank/internal/models"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) *fiber.App {
	controller := controller{db: db}
	return controller.route()
}

type controller struct {
	db *gorm.DB
}

func (c *controller) route() *fiber.App {
	app := fiber.New()
	{
		app.Get("/clients", c.Info())
		app.Get("/entity", c.EntityInfo())
		app.Get("/plan", c.PlanInfo())
		app.Get("/branch/:id", c.BranchInfo())
		app.Post("/upload", c.UploadData())
		app.Post("/login", c.Login())
		app.Get("/totals", c.Totals())
	}

	return app
}

func (c *controller) UploadData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("table")
		if err != nil {
			return ctx.Status(400).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		return ctx.SaveFile(file, fmt.Sprintf("uploads/%s", file.Filename))
	}
}

func (c *controller) Info() fiber.Handler {
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
			if err := rows.Scan(&rowData.Username, &rowData.CreatedAt, &rowData.OperationsCount, &rowData.OperationsPlus); err != nil {
				return ctx.Status(400).JSON(map[string]string{
					"error": err.Error(),
				})
			}

			res = append(res, &rowData)
		}

		return ctx.Status(200).JSON(res)
	}
}

func (c *controller) EntityInfo() fiber.Handler {
	type response struct {
		Count         int    `json:"count"`
		Type          string `json:"type"`
		IsLegalEntity bool   `json:"is_legal_entity"`
	}
	return func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Path())

		query := c.db.Table("operations").Select("count(operations.type), operations.type, clients.is_legal_entity").Joins(
			"left join clients on operations.client_id = clients.id")

		branchID := ctx.Query("branch_id")

		if branchID != "" {
			id, err := strconv.Atoi(branchID)
			if err != nil {
				return ctx.Status(400).JSON(map[string]string{
					"error": err.Error(),
				})
			}

			query = query.Where("branch_id = ?", id)
		}

		rows, err := query.Group("operations.type, clients.is_legal_entity").Rows()
		if err != nil {
			return ctx.Status(400).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		var res []*response

		for rows.Next() {
			var rowData response
			if err := rows.Scan(&rowData.Count, &rowData.Type, &rowData.IsLegalEntity); err != nil {
				return ctx.Status(400).JSON(map[string]string{
					"error": err.Error(),
				})
			}
			res = append(res, &rowData)
		}

		return ctx.Status(200).JSON(res)
	}
}

func (c *controller) PlanInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var queryResult []map[string]interface{}
		c.db.Table("employees").Select("employees.plan, count(employees.plan)").Group("employees.plan").Find(&queryResult)

		res := make(map[string]interface{}, 3)

		for _, v := range queryResult {
			key := fmt.Sprintf("%v", v["plan"])
			res[key] = v["count"].(int64)
		}

		return ctx.Status(200).JSON(res)
	}
}

func (c *controller) BranchInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.Status(400).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		var branch *models.Branch
		c.db.Find(&branch, &models.Branch{ID: uint(id)})
		return ctx.Status(200).JSON(branch)
	}
}

func (c *controller) Login() fiber.Handler {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Path())

		var req request
		ctx.BodyParser(&req)

		var manager models.Manager
		c.db.First(&manager, &models.Manager{Email: req.Email})

		if manager.ID == 0 {
			return ctx.Status(401).JSON(map[string]string{
				"error": "wrong user or password",
			})
		}

		if manager.Password != req.Password {
			return ctx.Status(401).JSON(map[string]string{
				"error": "wrong user or password",
			})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = manager.ID

		t, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
		if err != nil {
			return ctx.Status(400).JSON(map[string]string{
				"error": err.Error(),
			})
		}

		ctx.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    t,
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Hour * 24),
		})

		return ctx.Status(200).JSON(map[string]string{
			"success": "true",
		})
	}
}

func (c *controller) Totals() fiber.Handler {
	type response struct {
		Plus            int `json:"plus"`
		OperationsCount int `json:"operations_count"`
		CreditsCount    int `json:"credits_count"`

		LegalEntityPlus            int `json:"legal_entity_plus"`
		LegalEntityOperationsCount int `json:"legal_entity_operations_count"`
		LegalEntityCreditsCount    int `json:"legal_entity_credits_count"`
	}
	return func(ctx *fiber.Ctx) error {
		var res response

		// TODO: bad logic

		c.db.Table("operations").Select(
			"sum(amount)").Joins(
			"left join clients on operations.client_id = clients.id").Where(
			"clients.is_legal_entity = true").Scan(&res.Plus)
		c.db.Table("operations").Select("count(*)").Joins(
			"left join clients on operations.client_id = clients.id").Where(
			"clients.is_legal_entity = true").Scan(&res.OperationsCount)
		c.db.Table("operations").Select("count(*)").Joins(
			"left join clients on operations.client_id = clients.id").Where(
			"clients.is_legal_entity = true").Where("type = 'credit'").Scan(&res.CreditsCount)

		c.db.Table("operations").Select(
			"sum(amount)").Joins(
			"left join clients on operations.client_id = clients.id").Where(
			"clients.is_legal_entity = false").Scan(&res.LegalEntityPlus)
		c.db.Table("operations").Select("count(*)").Joins(
			"left join clients on operations.client_id = clients.id").Where(
			"clients.is_legal_entity = false").Scan(&res.LegalEntityOperationsCount)
		c.db.Table("operations").Select("count(*)").Joins(
			"left join clients on operations.client_id = clients.id").Where(
			"clients.is_legal_entity = false").Where("type = 'credit'").Scan(&res.LegalEntityCreditsCount)
		return ctx.Status(200).JSON(res)
	}
}
