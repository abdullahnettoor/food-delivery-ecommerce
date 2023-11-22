package entities

type CartItem struct {
	ID       uint `json:"-"`
	CartID   uint `json:"-"`
	DishID   uint `json:"-"`
	Quantity uint `json:"quantity"`
	Dish `gorm:"-"`

	FkDish Dish  `json:"-" gorm:"foreignkey:DishID;constraint:OnDelete:CASCADE"`
	FkCart  Cart `json:"-" gorm:"foreignkey:CartID;constraint:OnDelete:CASCADE"`
}

type Cart struct {
	ID        uint       `json:"cartId"`
	SellerID  uint       `json:"sellerId"`
	CartItems []CartItem `json:"dishes" gorm:"-"`

	FkSeller Seller `json:"-" gorm:"foreignkey:SellerID;constraint:OnDelete:CASCADE"`
}
