package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"os"

	"github.com/joho/godotenv"
)

type Product struct {
  gorm.Model
  Title string
  Slug string
  Price uint
}

func main() {
  err := godotenv.Load()
  if err != nil {
    panic("Error loading .env file")
  }

  dbHost := os.Getenv("DB_HOST")
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_Name")
  dbPort := os.Getenv("DB_PORT")

  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    dbHost, dbUser, dbPassword, dbName, dbPort,
  )
  
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err);
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{
    Title: "Airpods Pro 2nd Generation",
    Slug: "airpods-pro-2nd-generation",
    Price: 100,
  })

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}