package main

import (
	"ticket-engine/infrastructure/cache"
	"ticket-engine/infrastructure/datastore"
)

func main() {
	_ = datastore.NewDB()
	_ = cache.NewRedisClient()
}
