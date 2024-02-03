package main

import (
	"github.com/joho/godotenv"

	"ecommerce/api"

	"ecommerce/app"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	app.InitDB()

	api.SetupRoutes()
}
