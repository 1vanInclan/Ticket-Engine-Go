package datastore

import (
	"fmt"
	"log"
	"os"
	"ticket-engine/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewDB() *gorm.DB {

	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "ticket_user")
	password := getEnv("DB_PASSWORD", "ticket_password")
	dbname := getEnv("DB_NAME", "ticket_db")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

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
