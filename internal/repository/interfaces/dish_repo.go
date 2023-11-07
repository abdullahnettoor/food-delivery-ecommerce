package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type IDishRepository interface {
	FindPageWise(page int) (*[]entities.Dish, error)
	FindByID(id string) (*entities.Dish, error)
	Create(dish *entities.Dish) error
	Update(id string, dish *entities.Dish) (*entities.Dish, error)
	Delete(id string) error
}
