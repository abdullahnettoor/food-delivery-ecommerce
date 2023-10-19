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
	Status     UserStatus `json:"-" gorm:"default: Active"`
}
