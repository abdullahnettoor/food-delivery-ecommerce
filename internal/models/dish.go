package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model   `json:"-"`
	ID           uuid.UUID `json:"dishId" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RestaurantID uuid.UUID `json:"restaurantId" gorm:"foreignKey:restaurant.id;notNull"`
	Name         string    `json:"dishName" gorm:"notNull"`
	Description  string    `json:"dishDescription"`
	Price        float32   `json:"dishPrice" gorm:"notNull"`
	Quantity     uint      `json:"dishQuantity" gorm:"notNull"`
	Availability bool      `json:"isAvailable" gorm:"default:true"`
}
