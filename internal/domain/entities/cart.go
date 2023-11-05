package entities

type Cart struct {
	ID       uint `json:"cartId"`
	SellerID uint `json:"sellerId"`
}

type CartItem struct {
	ID       uint `json:"-"`
	DishID   uint `json:"dishId"`
	Quantity uint `json:"dishQuantiy"`
}
