package handlers

import (
	"ecommerce/models"
	"ecommerce/app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	var products []models.Product

	app.Db.Offset(offset).Limit(limit).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"limit":    limit,
		"products": products,
	})
}

func HandleSaveProduct(c *gin.Context) {
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app.Db.Create(&product)

	c.JSON(http.StatusCreated, product)
}
