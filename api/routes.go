package api

import (
	"github.com/gin-gonic/gin"
	"ecommerce/handlers"
)

func SetupRoutes() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	r.GET("/products", handlers.HandleGetProducts)
    r.POST("/products", handlers.HandleSaveProduct)
	r.POST("/auth/signup", handlers.HandleSignup)
	r.POST("/auth/login", handlers.HandleLogin)
	
	r.Run()
}
