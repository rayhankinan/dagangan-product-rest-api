package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"not null:unique"`
	Image       string `gorm:"default:null"`
	Description string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Stock       uint   `gorm:"default:0"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID;references:ID"`
}

type DisplayProduct struct {
	ID          uint      `json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Name        string    `json:"name,omitempty"`
	Image       string    `json:"image,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       uint      `json:"price,omitempty"`
	Stock       uint      `json:"stock,omitempty"`
	UserID      uint      `json:"owner_id,omitempty"`
}

func (product Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(&DisplayProduct{
		ID:          product.ID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		UserID:      product.UserID,
	})
}
