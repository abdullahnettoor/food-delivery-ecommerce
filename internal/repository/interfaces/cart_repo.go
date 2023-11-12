package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type ICartRepository interface {
	FindCart(id string) (*entities.Cart, error)
	FindCartItems(id string) (*[]entities.CartItem, error)
	AddToCart(id, dishId, sellerId string) error
	DecrementItem(id, dishId string) error
	DeleteItem(id, dishId string) error
}
