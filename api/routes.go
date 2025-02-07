package api

import (
	"ecommerce/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func SetupRoutes() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	r.Use(corsMiddleware())

	guardedRoutes := r.Group("/", authMiddleware)

	r.GET("/products", handlers.HandleGetProducts)
	r.GET("/products/:id", handlers.HandleGetProductById)

	r.POST("/auth/signup", handlers.HandleSignup)
	r.POST("/auth/login", handlers.HandleLogin)

	r.GET("/files/:fileName", handlers.HandleFetchFile)

	guardedRoutes.POST("/products", handlers.HandleSaveProduct)

	guardedRoutes.GET("/orders", handlers.HandleGetOrders)
	guardedRoutes.POST("/orders/init", handlers.HandleOrderInit)
	guardedRoutes.POST("/orders/success", handlers.HandleOrderSuccess)
	guardedRoutes.POST("/files/upload", handlers.HandleFilesUpload)

	guardedRoutes.POST("/address", handlers.HandleSaveAddress)
	guardedRoutes.GET("/address", handlers.HandleGetAddress)

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

func authMiddleware(c *gin.Context) {
	// Get the Authorization header value
	tokenString := c.GetHeader("Authorization")

	// Check if the token is missing
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized1"})
		c.Abort()
		return
	}

	secretKey := os.Getenv("JWT_SECRET")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Check for parsing errors
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized2"})
		c.Abort()
		return
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract userId from the token claims
		userId, ok := claims["userId"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId in token"})
			c.Abort()
			return
		}

		// Add userId to the Gin context for further use
		c.Set("userId", userId)

		// Continue with the next middleware or handler
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}
