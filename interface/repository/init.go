package repository

import (
	"ticket-engine/interface/repository/event"
	"ticket-engine/interface/repository/reservation"
	"ticket-engine/interface/repository/user"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	user.SetDB(db)
	event.SetDB(db)
	reservation.SetDB(db)
}
