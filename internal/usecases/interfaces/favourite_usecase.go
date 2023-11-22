package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type IFavouriteUseCase interface {
	AddFavItem(userId, dishId string) error
	ViewFavourites(userId string) (*[]entities.Dish, error)
	DeleteFavItem(userId, dishId string) error
}
