package event

import (
	"context"
	"errors"
	"ticket-engine/domain/model"
	usecaseRepo "ticket-engine/usecase/repository/event"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

var EventRepo usecaseRepo.EventRepository = &eventRepository{}

func SetDB(db *gorm.DB) {
	if repo, ok := EventRepo.(*eventRepository); ok {
		repo.db = db
	}
}

func (r *eventRepository) Create(ctx context.Context, event *model.Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *eventRepository) FindByID(ctx context.Context, id uint) (*model.Event, error) {
	var event model.Event
	err := r.db.WithContext(ctx).First(&event, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) FindAll(ctx context.Context) ([]*model.Event, error) {
	var events []*model.Event
	err := r.db.WithContext(ctx).Find(&events).Error
	return events, err
}

func (r *eventRepository) UpdateStock(ctx context.Context, eventID uint, newStock int) error {
	return r.db.WithContext(ctx).Model(&model.Event{}).
		Where("id = ?", eventID).
		Update("available_stock", newStock).Error
}
