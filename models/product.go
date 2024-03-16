package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Description string         `gorm:"not null" json:"description"`
	Price       int            `json:"price"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
