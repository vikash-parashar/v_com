package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `json:"product_name"`
	Description string    `json:"product_description"`
	Price       float64   `json:"product_price"`
	Image       string    `gorm:"type:text" json:"product_image"`
	Quantity    int       `json:"product_available_quantity"`
}
