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
	ShippingAddressID uint            `json:"-"`
	ShippingAddress   Address        `gorm:"foreignKey:ShippingAddressID" json:"shippingAddress"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
