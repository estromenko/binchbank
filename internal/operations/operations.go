package operations

import (
	"fmt"
	"os"
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
}
