package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"not null:unique"`
	Description string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Stock       uint   `gorm:"default:0"`
	Image       string `gorm:"default:null"`
	OwnerID     uint   `gorm:"not null"`
	Owner       User   `gorm:"foreignKey:OwnerID;references:ID"`
}

type DisplayProduct struct {
	ID          uint      `json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       uint      `json:"price,omitempty"`
	Stock       uint      `json:"stock,omitempty"`
	Image       string    `json:"image,omitempty"`
	OwnerID     uint      `json:"owner_id,omitempty"`
}

func (product Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(&DisplayProduct{
		ID:          product.ID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Image:       product.Image,
		OwnerID:     product.OwnerID,
	})
}
