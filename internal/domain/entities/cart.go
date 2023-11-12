package entities

type CartItem struct {
	ID       uint `json:"-"`
	CartID   uint `json:"-"`
	DishID   uint `json:"-"`
	Quantity uint `json:"quantity"`
	Dish     `gorm:"-"`
}

type Cart struct {
	ID        uint       `json:"cartId"`
	SellerID  uint       `json:"sellerId"`
	CartItems []CartItem `json:"dishes" gorm:"-"`
}
