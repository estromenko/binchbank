package main

import (
	"log"

	"github.com/estromenko/binchbank/internal/operations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	service := operations.New(db)
	service.Analyze()
}
