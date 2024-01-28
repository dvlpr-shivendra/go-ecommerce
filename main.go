package main

import (
	"fmt"
  "strconv"
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

  r := gin.Default()

  r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

  r.GET("/products", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))

		if page <= 0 {
			page = 1
		}

		if limit <= 0 {
			limit = 10
		}

		offset := (page - 1) * limit

		var products []Product
		db.Offset(offset).Limit(limit).Find(&products)

		c.JSON(http.StatusOK, gin.H{
			"page":       page,
			"limit":      limit,
			"products":   products,
		})
  })
  
  r.Run()
}