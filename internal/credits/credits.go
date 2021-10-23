package credits

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CreditService struct {
	db *gorm.DB
}

type Credit struct {
	ID     uint      `gorm:"primarykey" json:"id"`
	Type   string    `json:"type"`
	Amount uint      `json:"amount"`
	Date   time.Time `json:f"date"`
}

func (s *CreditService) Parse(file *os.File) ([]Credit, error) {
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

func (s *CreditService) Store() error {
	return nil
}

func (s *CreditService) Analyze() string {
	return ""
}

func New(db *gorm.DB) *CreditService {
	return &CreditService{
		db: db,
	}
}
