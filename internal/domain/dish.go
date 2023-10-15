package domain

import "github.com/google/uuid"

type Dish struct {
	ID           uuid.UUID `json:"dishId" gorm:"primaryKey"`
	RestaurantID uuid.UUID `json:"restaurantId"`
	Name         string    `json:"dishName"`
	Description  string    `json:"dishDescription"`
	Price        float32   `json:"dishPrice"`
	Quantity     uint      `json:"dishQuantity"`
	Availability bool      `json:"-" gorm:"default:true"`
}
