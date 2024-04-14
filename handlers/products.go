package handlers

import (
	"ecommerce/app"
	"ecommerce/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetProductById(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product

	if err := app.Db.First(&product, productID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func HandleGetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	var products []models.Product
	var total int64

	app.Db.Model(&models.Product{}).Count(&total)
	app.Db.Offset(offset).Limit(limit).Preload("Images").Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"limit":    limit,
		"total":    total,
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
