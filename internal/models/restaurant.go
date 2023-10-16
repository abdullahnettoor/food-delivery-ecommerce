package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestaurantStatus string

const (
	Verified RestaurantStatus = "Verified"
	Pending  RestaurantStatus = "Pending"
	Block    RestaurantStatus = "Blocked"
	Rejected RestaurantStatus = "Rejected"
)

type Restaurant struct {
	gorm.Model  `json:"-"`
	ID          uuid.UUID        `json:"restaurantId" gorm:"primaryKey" default:"gen_random_uuid()"`
	Name        string           `json:"restaurantName" gorm:"notNull"`
	Description string           `json:"restaurantDescription"`
	Email       string           `json:"email" gorm:"notNull"`
	Password    string           `json:"password" gorm:"notNull"`
	Status      RestaurantStatus `json:"-" gorm:"default: Pending"`
}
