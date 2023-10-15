package domain

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
	ID          uuid.UUID        `json:"restaurantId" gorm:"primaryKey"`
	Name        string           `json:"restaurantName"`
	Description string           `json:"restaurantDescription"`
	Email       string           `json:"email"`
	Password    string           `json:"password"`
	Status      RestaurantStatus `json:"-" gorm:"default: Pending"`
}
