package main

import (
	"ticket-engine/infrastructure/datastore"
)

func main() {
	_ = datastore.NewDB()
}
