package models

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"total_cart_price"`
	CreatedAt time.Time `json:"created_at"`
}
