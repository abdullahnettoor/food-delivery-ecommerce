package entities

import "time"

type OrderItem struct {
	ID       uint    `json:"-"`
	OrderID  uint    `json:"-"`
	DishID   uint    `json:"-"`
	Quantity uint    `json:"quantity"`
	// Price    float64 `json:"price"`
	SalePrice    float64 `json:"salePrice"`
	Dish     `gorm:"-"`
}

type Order struct {
	ID             uint        `json:"orderId"`
	UserID         uint        `json:"userId"`
	AddressID      uint        `json:"addressId"`
	SellerID       uint        `json:"sellerId"`
	OrderDate      time.Time   `json:"orderDate"`
	DeliveryDate   time.Time   `json:"deliveryDate"`
	PaymentMethod  string      `json:"paymentMethod"`
	TransactionID  string      `json:"transactionId"`
	Dishes         []OrderItem `json:"-" gorm:"-"`
	ItemCount      uint        `json:"itemCount"`
	Discount       float64     `json:"discount"`
	DeliveryCharge float64     `json:"deliveryCharge"`
	TotalPrice     float64     `json:"totalPrice"`
	Status         string      `json:"orderStatus"`
	PaymentStatus string      `json:"paymentStatus"`
}
