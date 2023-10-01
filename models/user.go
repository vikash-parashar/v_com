package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	Admin    = "admin"
	Customer = "customer"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	Orders    []Order   `gorm:"foreignKey:UserID" json:"orders"`
	Addresses []Address `json:"addresses"`
	CreatedAt time.Time `json:"created_at"`
}
