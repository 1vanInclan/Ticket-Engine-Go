package event

import (
	"context"
	"ticket-engine/domain/model"
)

type EventRepository interface {
	Create(ctx context.Context, event *model.Event) error
	FindByID(ctx context.Context, id uint) (*model.Event, error)
	FindAll(ctx context.Context) ([]*model.Event, error)
	UpdateStock(ctx context.Context, eventID uint, newStock int) error
}
