package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID             uuid.UUID `json:"orderId" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID         uuid.UUID `json:"userId" gorm:"type:uuid;foreignKey:users.id;notNull"`
	AddressID      uuid.UUID `json:"addressId" gorm:"type:uuid;foreignKey:addresses.id;notNull"`
	RestaurantID   uuid.UUID `json:"restaurantId" gorm:"foreignKey:restaurants.id;notNull"`
	PaymentMethod  string    `json:"paymentMethod" gorm:"default:COD"`
	TransactionID  uint      `json:"transactionId" gorm:"autoIncrement"`
	ItemCount      uint      `json:"itemCount"`
	Discount       float64   `json:"discount"`
	DeliveryCharge float64   `json:"deliveryCharge"`
	TotalPrice     float64   `json:"totalPrice"`
	Status         string    `json:"orderStatus" gorm:"default:Pending"`
	PayementStatus string    `json:"paymentStatus" gorm:"default:Pending"`
}

type OrderItem struct {
	ID       uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	OrderID  uuid.UUID `json:"orderId" gorm:"type:uuid;notNull"`
	DishID   uuid.UUID `json:"dishId" gorm:"type:uuid;notNull"`
	Quantity uint      `json:"dishQuantity"`
	Price    float64   `json:"dishPrice"`
}
