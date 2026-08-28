package main

import (
	"ticket-engine/infrastructure/cache"
	"ticket-engine/infrastructure/datastore"
	userRepo "ticket-engine/interface/repository/user"
)

func main() {
	db := datastore.NewDB()
	_ = cache.NewRedisClient()
	userRepo.SetDB(db)

}
