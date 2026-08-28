package reservation

import (
	"context"
	"errors"
	"ticket-engine/domain/model"
	usecaseRepo "ticket-engine/usecase/repository/reservation"

	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

var ReservationRepo usecaseRepo.ReservationRepository = &reservationRepository{}

func SetDB(db *gorm.DB) {
	if repo, ok := ReservationRepo.(*reservationRepository); ok {
		repo.db = db
	}
}

func (r *reservationRepository) Create(ctx context.Context, reservation *model.Reservation) error {
	return r.db.WithContext(ctx).Create(reservation).Error
}

func (r *reservationRepository) FindByID(ctx context.Context, id uint) (*model.Reservation, error) {
	var reservation model.Reservation
	err := r.db.WithContext(ctx).Preload("User").First(&reservation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &reservation, nil
}

func (r *reservationRepository) FindByUserID(ctx context.Context, userID uint) ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	err := r.db.WithContext(ctx).Preload("Event").Where("user_id = ?", userID).Find(&reservations).Error
	return reservations, err
}

func (r *reservationRepository) UpdateStatus(ctx context.Context, id uint, status model.ReservationStatus) error {
	return r.db.WithContext(ctx).Model(&model.Reservation{}).
		Where("id = ?", id).
		Update("status", status).Error
}
