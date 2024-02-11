package app

import (
	"fmt"

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
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	Db = db
	
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})

	if err != nil {
		panic(err)
	}
}
