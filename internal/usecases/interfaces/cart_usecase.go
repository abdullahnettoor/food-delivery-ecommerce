package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type ICartUseCase interface {
	AddtoCart(id, dishId string) error
	ViewCart(id string) (*entities.Cart, error)
	DeleteCartItem(id, dishId string) error
	DecrementCartItem(id, dishId string) error
	EmptyCart(id string) error
}
