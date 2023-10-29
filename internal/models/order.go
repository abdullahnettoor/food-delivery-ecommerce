package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Order struct {
	ID            uuid.UUID       `json:"cartId" gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID       `json:"userId" gorm:"type:uuid;foreignKey:user.id;notNull"`
	RestaurantID  uuid.UUID       `json:"restaurantId" gorm:"foreignKey:restaurants.id;notNull"`
	TransactionID string          `json:"transactionId"`
	Dishes        postgres.Hstore `json:"dishes" gorm:"type:hstore"`
	ItemCount     uint            `json:"itemCount"`
	TotalPrice    uint            `json:"totalPrice"`
	Status        string          `json:"orderStatus" gorm:"default:Pending"`
}
