package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model   `json:"-"`
	ID           uuid.UUID       `json:"cartId" gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID       `json:"restaurantId" gorm:"foreignKey:restaurants.id;notNull"`
	Dishes       postgres.Hstore `json:"dishes" gorm:"type:hstore"`
	ItemCount    uint            `json:"itemCount"`
	TotalPrice   uint            `json:"totalPrice"`
}
