package models

import (
	"time"
	"gorm.io/gorm"
)

type Address struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Line1      string         `gorm:"not null" json:"line1"`
	Line2      string         `json:"line2"`
	PostalCode string         `json:"postalCode"`
	Landmark   string         `json:"landmark"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Phone     string         `gorm:"uniqueIndex;size:15" json:"phone"`
	Email     string         `gorm:"uniqueIndex;size:255" json:"email"`
	Password  string         `gorm:"size:1000" json:"-"`
	AddressId int            `gorm:"default:null" json:"-"`
	Address   *Address       `json:"address,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}