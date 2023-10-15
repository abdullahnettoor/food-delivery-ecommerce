package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStatus string

const (
	Active  UserStatus = "Active"
	Blocked UserStatus = "Blocked"
	Deleted UserStatus = "Deleted"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID  `json:"userId" gorm:"primaryKey"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email" gorm:"notNull"`
	Phone      string     `json:"phone"`
	Password   string     `json:"password"`
	Status     UserStatus `json:"-" gorm:"default: Active"`
}
