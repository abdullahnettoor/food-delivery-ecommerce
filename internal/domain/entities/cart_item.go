package entities

type CartItem struct {
	ID       uint `json:"-"`
	DishID   uint `json:"dishId"`
	Quantity uint `json:"dishQuantiy"`
}
