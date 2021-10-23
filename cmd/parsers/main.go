package main

import (
	"fmt"
	"log"
	"os"

	"github.com/estromenko/binchbank/internal/operations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Open("data/test.csv")

	db.AutoMigrate(&operations.Operation{})

	data, err := operations.Analyze(db, file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
