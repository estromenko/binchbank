package operations

import (
	"fmt"
	"os"
<<<<<<< HEAD
	"regexp"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type OperationService struct {
	db *gorm.DB
}

type Operation struct {
	ID   uint      `gorm:"primarykey" mapstructure:"id"`
	Date time.Time `mapstructure:"date"`
}

func Parse(file *os.File) error {

	match, err := regexp.MatchString(`.+\.(json|yaml|yml)$`, file.Name())
	if err != nil {
		return err
	}

	if !match {
		return fmt.Errorf("unsupported file type")
	}

	var operation Operation

	if err := viper.ReadConfig(file); err != nil {
		return err
	}
	return viper.Unmarshal(&operation)
}

func (s *OperationService) Store() error {
	return nil
}

func (s *OperationService) Analyze() string {
	return ""
}

func New(db *gorm.DB) *OperationService {
	return &OperationService{
		db: db,
	}
=======
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
>>>>>>> 03b2f2fd2bc76b58d6ae3aea5c3694b959a8bac3
}
