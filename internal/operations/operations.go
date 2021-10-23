package operations

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Operation struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Date     time.Time `json:"date"`
	Username string    `json:"username"`
}

func parse(file *os.File) ([]Operation, error) {
	nameParts := strings.Split(file.Name(), ".")

	if len(nameParts) <= 1 {
		return nil, fmt.Errorf("invalid file extension")
	}

	ext := nameParts[len(nameParts)-1]

	if ext == "json" {
		return parseJSON(file)
	}

	if ext == "csv" {
		return parseCSV(file)
	}

	return nil, fmt.Errorf("unsupported file path")
}

func Analyze(db *gorm.DB, file *os.File) (string, error) {
	operations, err := parse(file)
	if err != nil {
		return "", err
	}

	for _, operation := range operations {
		db.Create(&operation)
	}

	var amount int64
	db.Find(&Operation{}).Count(&amount)

	return string(rune(amount)), nil
}
