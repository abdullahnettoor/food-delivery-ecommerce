package models

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
	ID         uuid.UUID  `json:"userId" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName  string     `json:"firstName" gorm:"notNull"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email" gorm:"notNull"`
	Phone      string     `json:"phone" gorm:"notNull"`
	Password   string     `json:"-" gorm:"notNull"`
	Status     UserStatus `json:"status" gorm:"default: Active"`
}

type Address struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"addressId" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID `json:"userId" gorm:"type:uuid;foreignKey:user.id;notNull"`
	City       string    `json:"street" gorm:"notNull"`
	District   string    `json:"district" gorm:"notNull"`
	State      string    `json:"state" gorm:"notNull"`
	PinCode    string    `json:"pinCode" gorm:"notNull"`
}
