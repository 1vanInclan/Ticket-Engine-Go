package main

import (
	"ticket-engine/infrastructure/cache"
	"ticket-engine/infrastructure/datastore"
	"ticket-engine/interface/repository"
)

func main() {
	db := datastore.NewDB()
	_ = cache.NewRedisClient()

	// Inicializar repositorios
	repository.Init(db)

}
