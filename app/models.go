package app

import "time"

type Operation struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Date     time.Time `json:"date"`
	Username string    `json:"username"`
}

type Credit struct {
	ID     uint      `gorm:"primarykey" json:"id"`
	Type   string    `json:"type"`
	Amount uint      `json:"amount"`
	Date   time.Time `json:"date"`
}
