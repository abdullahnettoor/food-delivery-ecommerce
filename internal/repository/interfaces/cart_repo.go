package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type ICartRepository interface {
	CreateCart(id, sellerId string) error
	FindCart(id string) (*entities.Cart, error)
	FindCartItems(id string) (*[]entities.CartItem, error)
	AddToCart(id, dishId string) error
	DecrementItem(id, dishId string) error
	DeleteItem(id, dishId string) error
	DeleteCart(id string) error
}
