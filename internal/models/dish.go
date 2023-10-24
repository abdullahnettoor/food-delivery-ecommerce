package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model   `json:"-"`
	ID           uuid.UUID `json:"dishId" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RestaurantID uuid.UUID `json:"restaurantId" gorm:"type:uuid;foreignKey:restaurant.id;notNull"`
	Name         string    `json:"dishName" gorm:"notNull"`
	Description  string    `json:"dishDescription"`
	Price        float32   `json:"dishPrice" gorm:"notNull"`
	Quantity     uint      `json:"dishQuantity" gorm:"notNull"`
	Category     uint      `json:"category" gorm:"type:bigInt;foreignKey:category.id;notNull"`
	IsVeg        bool      `json:"isVeg" gorm:"default:false"`
	Availability bool      `json:"isAvailable" gorm:"default:true"`
}
