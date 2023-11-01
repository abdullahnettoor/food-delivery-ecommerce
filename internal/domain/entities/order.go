package entities

type Order struct {
	ID             uint    `json:"orderId"`
	UserID         uint    `json:"userId"`
	AddressID      uint    `json:"addressId"`
	RestaurantID   uint    `json:"restaurantId"`
	PaymentMethod  string  `json:"paymentMethod"`
	TransactionID  uint    `json:"transactionId"`
	ItemCount      uint    `json:"itemCount"`
	Discount       float64 `json:"discount"`
	DeliveryCharge float64 `json:"deliveryCharge"`
	TotalPrice     float64 `json:"totalPrice"`
	Status         string  `json:"orderStatus"`
	PayementStatus string  `json:"paymentStatus"`
}
