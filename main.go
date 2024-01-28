package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"os"

	"github.com/joho/godotenv"
)

type Product struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Price     uint           `json:"price"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
		panic(err)
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
			"page":     page,
			"limit":    limit,
			"products": products,
		})
	})

	r.POST("/products", func(c *gin.Context) {
		var product Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&product)

		c.JSON(http.StatusCreated, product)
	})

	r.Run()
}
