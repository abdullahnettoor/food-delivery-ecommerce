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
	ID          uuid.UUID        `json:"restaurantId" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string           `json:"restaurantName" gorm:"notNull"`
	Description string           `json:"restaurantDescription"`
	Email       string           `json:"email" gorm:"notNull;unique"`
	Password    string           `json:"-" gorm:"notNull"`
	Status      RestaurantStatus `json:"status" gorm:"default: Pending"`
}
