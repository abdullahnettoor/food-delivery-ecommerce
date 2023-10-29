package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	ID           uuid.UUID `json:"cartId" gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID `json:"restaurantId" gorm:"type:uuid"`
}

type CartItem struct {
	ID       uint      `json:"-" gorm:"type:bigInt;primaryKey;autoIncrement"`
	CartID   uuid.UUID `json:"cartId" gorm:"type:uuid;notNull"`
	DishID   uuid.UUID `json:"dishId" gorm:"type:uuid;"`
	Quantity uint      `json:"dishQuantity" gorm:"type:uint"`
}
