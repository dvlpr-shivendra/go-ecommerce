package handlers

import (
	"ecommerce/app"
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSaveAddress(c *gin.Context) {
	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address.UserId = uint(c.GetFloat64("userId"))

	app.Db.Create(&address)

	c.JSON(http.StatusOK, address)
}

func HandleGetAddress(c *gin.Context) {
	var addresses []models.Address
	app.Db.Where("user_id = ?", uint(c.GetFloat64("userId"))).Find(&addresses)
	c.JSON(http.StatusOK, addresses)
}
