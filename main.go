package main

import (
	"ticket-engine/infrastructure/cache"
	"ticket-engine/infrastructure/datastore"
	eventRepo "ticket-engine/interface/repository/event"
	reservationRepo "ticket-engine/interface/repository/reservation"
	userRepo "ticket-engine/interface/repository/user"
)

func main() {
	db := datastore.NewDB()
	_ = cache.NewRedisClient()

	userRepo.SetDB(db)
	eventRepo.SetDB(db)
	reservationRepo.SetDB(db)

}
