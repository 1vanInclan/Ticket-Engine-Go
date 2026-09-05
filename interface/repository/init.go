package repository

import (
	"ticket-engine/interface/repository/event"
	"ticket-engine/interface/repository/reservation"
	"ticket-engine/interface/repository/user"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, rdb *redis.Client) {
	user.SetDB(db)

	event.SetDB(db)
	event.SetRedis(rdb)

	reservation.SetDB(db)
	reservation.SetRedis(rdb)

}
