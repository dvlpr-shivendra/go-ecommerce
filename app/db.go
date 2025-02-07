package app

import (
	"fmt"
	"log"
	"time"

	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	"ecommerce/models"
)

var Db *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 3; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	Db = db

	var orderStatusExists bool

	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status')").Scan(&orderStatusExists)
	
	if !orderStatusExists {
		query := "CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped', 'delivered', 'canceled')"
		if err := db.Exec(query).Error; err != nil {
			log.Printf("Error creating order_status enum: %v\n", err)
		}
	}

	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})

	if err != nil {
		panic(err)
	}
}
