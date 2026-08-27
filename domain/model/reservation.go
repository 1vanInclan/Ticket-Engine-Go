package model

import "time"

type ReservationStatus string

const (
	StatusPending   ReservationStatus = "PENDING"
	StatusConfirmed ReservationStatus = "CONFIRMED"
	StatusExpired   ReservationStatus = "EXPIRED"
	StatusCancelled ReservationStatus = "CANCELLED"
)

type Reservation struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	UserID    uint              `gorm:"not null;index" json:"user_id"`
	EventID   uint              `gorm:"not null;index" json:"event_id"`
	Quantity  int               `gorm:"not null" json:"quantity"`
	Status    ReservationStatus `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`

	// Relaciones GORM
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Event Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}
