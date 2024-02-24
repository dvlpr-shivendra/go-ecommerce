package handlers

import (
	"ecommerce/app"
	"net/http"
	"os"

	"ecommerce/models"

	"ecommerce/helpers"

	"github.com/gin-gonic/gin"

	razorpay "github.com/razorpay/razorpay-go"
)

type OrderInitRequest struct {
	ProductId    string `json:"productId"`
}

func HandleOrderInit(c *gin.Context) {
	key := os.Getenv("RAZORPAY_KEY")
	secret := os.Getenv("RAZORPAY_SECRET")
	client := razorpay.NewClient(key, secret)

	var orderInitRequest OrderInitRequest

	if err := c.BindJSON(&orderInitRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productId, err := helpers.StrToUint(orderInitRequest.ProductId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad product id"})
		return
	}

	var product models.Product
	var order models.Order

	if err := app.Db.First(&product, productId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	order.ProductId = productId
	order.UserId = 1
	order.Status = "pending"
	order.ShippingAddressID = 1
	order.BillingAddressID = 1
	order.Amount = product.Price
	order.AmountDue = product.Price

	// Create a new order
	params := map[string]interface{}{
		"amount":          product.Price * 100, // Amount in paise (e.g., 1000 paise = ₹10.00)
		"currency":        "INR",
		"receipt":         "order_rcptid_11",
		"payment_capture": 1,
	}

	razorpayOrder, err := client.Order.Create(params, nil)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating order"})
		return
	}

	createProduct := app.Db.Create(&order)

	if createProduct.Error != nil {
		c.JSON(500, gin.H{"error": "Error creating order"})
		return
	}

	// Return order ID to the client
	c.JSON(200, razorpayOrder)
}

func HandleOrderSuccess(c *gin.Context) {

}
