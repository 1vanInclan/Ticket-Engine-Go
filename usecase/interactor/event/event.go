package event

import (
	"context"
	eventImpl "ticket-engine/interface/repository/event"
	"ticket-engine/usecase/dto"
	eventRepo "ticket-engine/usecase/repository/event"
)

type EventInteractor interface {
	Create(ctx context.Context, input dto.CreateEventInput) (*dto.EventResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.EventResponse, error)
	FindAll(ctx context.Context) ([]*dto.EventResponse, error)
}

type eventInteractor struct {
	eventRepository eventRepo.EventRepository
}

var EventInt EventInteractor = &eventInteractor{
	eventRepository: eventImpl.EventRepo,
}
