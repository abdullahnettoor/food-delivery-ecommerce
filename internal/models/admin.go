package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"adminId" gorm:"primaryKey" default:"gen_random_uuid()"`
	FirstName  string    `json:"firstName" gorm:"notNull"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email" gorm:"notNull"`
	Password   string    `json:"-" gorm:"notNull"`

}
