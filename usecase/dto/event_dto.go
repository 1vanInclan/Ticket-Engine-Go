package dto

import "time"

type CreateEventInput struct {
	Name          string  `json:"name"`
	TotalCapacity int     `json:"total_capacity"`
	Price         float64 `json:"price"`
}

type EventResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	TotalCapacity  int       `json:"total_capacity"`
	AvailableStock int       `json:"available_stock"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
