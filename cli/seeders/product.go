package main

import (
	"ecommerce/app"
	"ecommerce/models"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
)

// Usage: go run cli/seeders/product.go <count>
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	app.InitDB()

	count, err := strconv.Atoi(os.Args[1])
	
	if err != nil {
		panic("Please provide a count as argument")
	}

	for i := 0; i < count; i++ {
		product := models.Product{
			Title:       faker.Sentence(),
			Images:      []string{
				"https://unsplash.it/20" + strconv.Itoa(i),
				"https://unsplash.it/20" + strconv.Itoa(i+1),
				"https://unsplash.it/20" + strconv.Itoa(i+2),
			},
			Description: faker.Paragraph(),
			Price:       100 + rand.Intn(1000-100),
		}
		app.Db.Create(&product)
	}
}
