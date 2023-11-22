package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type IFavoriteRepository interface {
	Create(userId, dishId string, item *entities.Favourite) error
	FindByUserId(userId string) (*[]entities.Favourite, error)
	Delete(userId, dishId string) error
}
