package handlers

import (
	"os"

	"net/http"

	"ecommerce/app"

	"ecommerce/models"

	"github.com/gin-gonic/gin"

	razorpay "github.com/razorpay/razorpay-go"
)

type OrderInitRequest struct {
	ProductId uint `json:"productId"`
}

type OrderSuccessRequest struct {
	OrderId           uint   `json:"orderId"`
	RazorpayPaymentId string `json:"razorpayPaymentId"`
}

func getRazorPayClinet() *razorpay.Client {
	key := os.Getenv("RAZORPAY_KEY")
	secret := os.Getenv("RAZORPAY_SECRET")
	return razorpay.NewClient(key, secret)
}

func HandleOrderInit(c *gin.Context) {
	client := getRazorPayClinet()

	var orderInitRequest OrderInitRequest

	if err := c.BindJSON(&orderInitRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	var order models.Order

	if err := app.Db.First(&product, orderInitRequest.ProductId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	order.ProductId = product.ID
	order.UserId = uint(c.GetFloat64("userId"))
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order"})
		return
	}

	if err := app.Db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order"})
		return
	}

	// Return order ID to the client
	c.JSON(200, gin.H{
		"razorpayOrder": razorpayOrder,
		"order":         order,
	})
}

func HandleOrderSuccess(c *gin.Context) {
	var request OrderSuccessRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order

	if err := app.Db.First(&order, request.OrderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	client := getRazorPayClinet()

	payment, err := client.Payment.Fetch(request.RazorpayPaymentId, nil, nil)

	if err != nil {
		if err := app.Db.Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
	}

	receivedAmount := int(payment["amount"].(float64))

	if payment["status"].(string) != "authorized" || receivedAmount != order.Amount * 100 {

		// TODO: can we return the user's money here??

		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

		return
	}

	_, err = client.Payment.Capture(request.RazorpayPaymentId, receivedAmount, nil, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	order.Status = "processing"
	order.AmountDue = 0

	if err := app.Db.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(200, gin.H{"order": order})
}
