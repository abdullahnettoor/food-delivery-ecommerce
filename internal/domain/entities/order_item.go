package entities

type OrderItem struct {
	ID       uint    `json:"-"`
	OrderID  uint    `json:"orderId"`
	DishID   uint    `json:"dishId"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}
