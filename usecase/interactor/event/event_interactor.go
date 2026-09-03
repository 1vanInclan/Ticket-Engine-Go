package event

import (
	"context"
	"errors"
	"ticket-engine/domain/model"
	"ticket-engine/usecase/dto"
)

func (i *eventInteractor) Create(ctx context.Context, input dto.CreateEventInput) (*dto.EventResponse, error) {
	if input.Name == "" {
		return nil, errors.New("event name is required")
	}
	if input.TotalCapacity <= 0 {
		return nil, errors.New("total capacity must be greater than zero")
	}

	newEvent := &model.Event{
		Name:           input.Name,
		TotalCapacity:  input.TotalCapacity,
		AvailableStock: input.TotalCapacity,
		Price:          input.Price,
	}

	if err := i.eventRepository.Create(ctx, newEvent); err != nil {
		return nil, err
	}

	if err := i.eventRepository.SetStock(ctx, newEvent.ID, newEvent.TotalCapacity); err != nil {
		return nil, errors.New("error al inicializar stock en redis")
	}

	return toEventResponse(newEvent), nil
}

func (i *eventInteractor) FindByID(ctx context.Context, id uint) (*dto.EventResponse, error) {
	event, err := i.eventRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event not found")
	}

	return toEventResponse(event), nil
}

func (i *eventInteractor) FindAll(ctx context.Context) ([]*dto.EventResponse, error) {
	events, err := i.eventRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []*dto.EventResponse
	for _, e := range events {
		response = append(response, toEventResponse(e))
	}

	return response, nil
}

func toEventResponse(e *model.Event) *dto.EventResponse {
	return &dto.EventResponse{
		ID:             e.ID,
		Name:           e.Name,
		TotalCapacity:  e.TotalCapacity,
		AvailableStock: e.AvailableStock,
		Price:          e.Price,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}
