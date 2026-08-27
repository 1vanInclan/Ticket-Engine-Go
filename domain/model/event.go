package model

import "time"

type Event struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	TotalCapacity  int       `gorm:"not null" json:"total_capacity"`
	AvailableStock int       `gorm:"not null" json:"available_stock"`
	Price          float64   `gorm:"not null" json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
