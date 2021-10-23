package clients

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
}

type OperationType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Operation struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	ClientID      int           `json:"client_id"`
	Client        Client        `gorm:"foreignKey:"ClientID"`
	Type          int           `json:"type"`
	OperationType OperationType `gorm:"foreignKey:Type"`
	Date          time.Time     `json:"date"`
}
