package reservation

import (
	"context"
	"errors"
	"fmt"
	"ticket-engine/domain/model"
	usecaseRepo "ticket-engine/usecase/repository/reservation"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

var ReservationRepo usecaseRepo.ReservationRepository = &reservationRepository{}

func SetDB(db *gorm.DB) {
	if repo, ok := ReservationRepo.(*reservationRepository); ok {
		repo.db = db
	}
}

func SetRedis(rbd *redis.Client) {
	if repo, ok := ReservationRepo.(*reservationRepository); ok {
		repo.rdb = rbd
	}
}

// Operaciones Atómicas en Redis
func (r *reservationRepository) DecrementStock(ctx context.Context, eventID uint, quantity int) (int64, error) {
	key := fmt.Sprintf("event:stock:%d", eventID)
	// DECRBY resta la cantidad especificada de forma atómica
	return r.rdb.DecrBy(ctx, key, int64(quantity)).Result()
}

func (r *reservationRepository) IncrementStock(ctx context.Context, eventID uint, quantity int) (int64, error) {
	key := fmt.Sprintf("event:stock:%d", eventID)
	// INCRBY devuelve el stock si una reserva fue cancelada o rechazada por falta de cupo
	return r.rdb.IncrBy(ctx, key, int64(quantity)).Result()
}

// Persistencia en PostgreSQL
func (r *reservationRepository) Create(ctx context.Context, reservation *model.Reservation) error {
	return r.db.WithContext(ctx).Create(reservation).Error
}

func (r *reservationRepository) FindByID(ctx context.Context, id uint) (*model.Reservation, error) {
	var reservation model.Reservation
	err := r.db.WithContext(ctx).Preload("Event").First(&reservation, id).Error
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
