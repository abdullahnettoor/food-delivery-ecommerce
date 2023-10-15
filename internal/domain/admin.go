package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"adminId" gorm:"primaryKey"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Status     bool      `json:"-" gorm:"default: true"`
}
