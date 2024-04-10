package models

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Line1      string         `gorm:"not null" json:"line1"`
	Line2      string         `json:"line2"`
	PostalCode string         `json:"postalCode"`
	Landmark   string         `json:"landmark"`
	UserId     uint           `json:"-"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
