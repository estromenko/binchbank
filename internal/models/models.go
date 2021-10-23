package models

import (
	"time"

	"gorm.io/gorm"
)

type Operation struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Date     time.Time `json:"date"`
	ClientID uint      `json:"client_id"`
	Client   Client    `gorm:"foreignKey:ClientID"`
	Amount   int       `json:"amount"`
	Type     string    `json:"type"`
	BranchID uint      `json:"branch_id"`
	Branch   Branch    `gorm:"foreignKey:BranchID"`
}

type Client struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Username      string    `json:"username"`
	CreatedAt     time.Time `json:"created_at"`
	IsLegalEntity bool      `json:"is_legal_entity"`
}

type Branch struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Rating      uint   `json:"rating"`
	ManagerName string `json:"manager_name"`
}

type Employee struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	RoleID    uint      `json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	Plan      uint      `json:"plan"`
	BranchID  uint      `json:"branch_id"`
	Branch    Branch    `gorm:"foreignKey:BranchID"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Manager struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	IsTop    bool   `json:"is_top"`
	Password string `json:"password"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Operation{},
		&Client{},
		&Branch{},
		&Employee{},
		&Role{},
		&Manager{},
	)
}
