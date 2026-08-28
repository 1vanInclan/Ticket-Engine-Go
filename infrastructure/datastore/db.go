package datastore

import (
	"fmt"
	"log"
	"ticket-engine/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=ticket_user password=ticket_password dbname=ticket_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established succesfully")

	// Automigracion de las entidades de dominio
	err = db.AutoMigrate(&model.User{}, &model.Event{}, &model.Reservation{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	fmt.Println("Database migration completed")

	return db

}
