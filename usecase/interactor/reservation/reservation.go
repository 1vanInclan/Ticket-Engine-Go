package reservation

import (
	"context"
	eventImpl "ticket-engine/interface/repository/event"
	resImpl "ticket-engine/interface/repository/reservation"
	"ticket-engine/usecase/dto"
	eventRepo "ticket-engine/usecase/repository/event"
	resRepo "ticket-engine/usecase/repository/reservation"
)

type ReservationInteractor interface {
	Create(ctx context.Context, userID uint, input dto.CreateReservationInput) (*dto.ReservationResponse, error)
	FindByUserID(ctx context.Context, userID uint) ([]*dto.ReservationResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.ReservationResponse, error)
}

type reservationInteractor struct {
	reservationRepo resRepo.ReservationRepository
	eventRepo       eventRepo.EventRepository
}

var ReservationInt ReservationInteractor = &reservationInteractor{
	reservationRepo: resImpl.ReservationRepo,
	eventRepo:       eventImpl.EventRepo,
}
