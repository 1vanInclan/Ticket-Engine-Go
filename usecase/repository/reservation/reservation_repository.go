package reservation

import (
	"context"
	"ticket-engine/domain/model"
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation *model.Reservation) error
	FindByID(ctx context.Context, id uint) (*model.Reservation, error)
	FindByUserID(ctx context.Context, userID uint) ([]*model.Reservation, error)
	UpdateStatus(ctx context.Context, id uint, status model.ReservationStatus) error
}
