package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	ID           uuid.UUID          `json:"cartId"`
	RestaurantID uuid.UUID          `json:"restaurantId"`
	Dishes       map[uuid.UUID]uint `json:"dishes"`
	TotalPrice   float64            `json:"totalPrice"`
}

type CartItem struct {
	ID           uuid.UUID `gorm:"type:uuid;notNull"`
	RestaurantID uuid.UUID `gorm:"type:uuid"`
	DishID       uuid.UUID `gorm:"type:uuid;unique"`
	Quantity     uint      `gorm:"type:uint"`
}
