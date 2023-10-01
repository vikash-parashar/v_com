package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	Resident  = "resident"
	Permanent = "permanent"
	Current   = "current"
)

type Address struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Type      string    `json:"address_type"`
	CreatedAt time.Time `json:"created_at"`
}
