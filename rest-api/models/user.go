package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"

	"dagangan-product-rest-api/types"
)

type User struct {
	gorm.Model
	Username       string             `gorm:"not null;unique"`
	HashedPassword types.HashedString `gorm:"not null"`
	Role           types.AuthRole     `gorm:"not null"`
	Products       []Product
}

type DisplayUser struct {
	ID             uint               `json:"id,omitempty"`
	CreatedAt      time.Time          `json:"created_at,omitempty"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty"`
	Username       string             `json:"username,omitempty"`
	HashedPassword types.HashedString `json:"-"`
	Role           types.AuthRole     `json:"role,omitempty"`
	Products       []Product          `json:"managed_product,omitempty"`
}

func (user User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&DisplayUser{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Role:      user.Role,
		Products:  user.Products,
	})
}
