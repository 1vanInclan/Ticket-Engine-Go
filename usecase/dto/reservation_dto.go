package dto

import (
	"ticket-engine/domain/model"
	"time"
)

type CreateReservationInput struct {
	EventID  uint `json:"event_id"`
	Quantity int  `json:"quantity"`
}

type ReservationResponse struct {
	ID        uint                    `json:"id"`
	UserID    uint                    `json:"user_id"`
	EventID   uint                    `json:"event_id"`
	Quantity  int                     `json:"quantity"`
	Status    model.ReservationStatus `json:"status"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	Event     *EventResponse          `json:"event,omitempty"`
}
