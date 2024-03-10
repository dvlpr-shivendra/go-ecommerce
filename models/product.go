package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ProductId uint           `json:"-"`
	Url       string         `json:"url"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Images      []ProductImage `json:"images"`
	Description string         `gorm:"not null" json:"description"`
	Price       int            `json:"price"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
