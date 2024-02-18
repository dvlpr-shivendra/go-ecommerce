package handlers

import (
	"ecommerce/models"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"ecommerce/app"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleSignup(c *gin.Context) {
	var signupRequest SignupRequest

	if err := c.BindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user.Name = signupRequest.Name
	user.Email = signupRequest.Email
	user.Phone = signupRequest.Phone
	user.Password = string(password)

	result := app.Db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": result.Error})
		return
	}

	tokenStr, err := generateJWT(user.ID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenStr,
		"user": user,
	})
}

func HandleLogin(c *gin.Context) {
	var user models.User
	var loginRequest LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := app.Db.Where("email = ?", loginRequest.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	tokenStr, err := generateJWT(user.ID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusContinue, gin.H{
		"token": tokenStr,
		"user": user,
	})
}

func generateJWT(userId uint) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userId":    userId,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}
