package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	Cancelled = "cancelled"
	Delivered = "delivered"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Products  []Product `gorm:"many2many:order_products" json:"products"`
	Price     float64   `json:"total_order_price"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"order_status"`
}

// Define the many-to-many relationship between Order and Product
type OrderProduct struct {
	OrderID   uuid.UUID `gorm:"type:uuid"`
	ProductID uuid.UUID `gorm:"type:uuid"`
}
