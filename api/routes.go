package api

import (
	"ecommerce/handlers"

	"github.com/gin-gonic/gin"
	
)

func SetupRoutes() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	r.Use(corsMiddleware())

	r.POST("/products", handlers.HandleSaveProduct)
	r.GET("/products", handlers.HandleGetProducts)
	r.GET("/products/:id", handlers.HandleGetProductById)

	r.POST("/auth/signup", handlers.HandleSignup)
	r.POST("/auth/login", handlers.HandleLogin)
	r.POST("/order/init", handlers.HandleOrderInit)
	r.POST("/order/success", handlers.HandleOrderSuccess)

	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
