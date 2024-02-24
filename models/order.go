package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Delivered  OrderStatus = "delivered"
	Canceled   OrderStatus = "canceled"
)

type Order struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	Status            OrderStatus    `gorm:"type:order_status" json:"status"`
	ShippingAddressID uint           `json:"-"`
	ShippingAddress   Address        `gorm:"foreignKey:ShippingAddressID" json:"shippingAddress"`
	BillingAddressID  uint           `json:"-"`
	BillingAddress    Address        `gorm:"foreignKey:BillingAddressID" json:"billingAddress"`
	UserId            uint           `json:"-"`
	Product           Product        `json:"product"`
	ProductId         uint           `json:"-"`
	Amount            uint           `json:"amount"`
	AmountDue         uint           `json:"amountDue"`
	User              User           `json:"user"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
