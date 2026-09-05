package reservation

import (
	"context"
	"errors"
	"ticket-engine/domain/model"
	"ticket-engine/usecase/dto"
)

func (i *reservationInteractor) Create(ctx context.Context, userID uint, input dto.CreateReservationInput) (*dto.ReservationResponse, error) {
	if input.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	// 1. Validar que el evento exista en PostgreSQL
	event, err := i.eventRepo.FindByID(ctx, input.EventID)
	if err != nil {
		return nil, err
	}
	if event == nil {
		return nil, errors.New("event not found")
	}

	// 2. Control de Concurrencia Atómico en Redis
	newStock, err := i.reservationRepo.DecrementStock(ctx, input.EventID, input.Quantity)
	if err != nil {
		return nil, errors.New("failed to process stock operation in redis")
	}

	// 3. Rollback inmediato en Redis si el stock es insuficiente
	if newStock < 0 {
		_, _ = i.reservationRepo.IncrementStock(ctx, input.EventID, input.Quantity)
		return nil, errors.New("insufficient stock for this event")
	}

	// 4. Crear la reservación en PostgreSQL con estado PENDING
	newReservation := &model.Reservation{
		UserID:   userID,
		EventID:  input.EventID,
		Quantity: input.Quantity,
		Status:   model.StatusPending,
	}

	if err := i.reservationRepo.Create(ctx, newReservation); err != nil {
		// Rollback del stock en Redis si falla la escritura en Postgres
		_, _ = i.reservationRepo.IncrementStock(ctx, input.EventID, input.Quantity)
		return nil, errors.New("failed to create reservation record")
	}

	newReservation.Event = *event
	return toReservationResponse(newReservation), nil
}

func (i *reservationInteractor) FindByUserID(ctx context.Context, userID uint) ([]*dto.ReservationResponse, error) {
	reservations, err := i.reservationRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.ReservationResponse, 0, len(reservations))
	for _, res := range reservations {
		responses = append(responses, toReservationResponse(res))
	}

	return responses, nil
}

func (i *reservationInteractor) FindByID(ctx context.Context, id uint) (*dto.ReservationResponse, error) {
	reservation, err := i.reservationRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if reservation == nil {
		return nil, nil
	}

	return toReservationResponse(reservation), nil
}

func toReservationResponse(r *model.Reservation) *dto.ReservationResponse {
	resp := &dto.ReservationResponse{
		ID:        r.ID,
		UserID:    r.UserID,
		EventID:   r.EventID,
		Quantity:  r.Quantity,
		Status:    r.Status,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	if r.Event.ID != 0 {
		resp.Event = &dto.EventResponse{
			ID:             r.Event.ID,
			Name:           r.Event.Name,
			TotalCapacity:  r.Event.TotalCapacity,
			AvailableStock: r.Event.AvailableStock,
			Price:          r.Event.Price,
			CreatedAt:      r.Event.CreatedAt,
			UpdatedAt:      r.Event.UpdatedAt,
		}
	}

	return resp
}
