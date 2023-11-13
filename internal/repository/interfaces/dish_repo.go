package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type IDishRepository interface {
	Search(search string) (*[]entities.Dish, error)
	FindPageWise(page, limit uint) (*[]entities.Dish, error)
	FindByID(id string) (*entities.Dish, error)
	FindBySeller(sellerId string) (*[]entities.Dish, error)
	FindBySellerAndID(id, sellerId string) (*entities.Dish, error)
	Create(dish *entities.Dish) error
	Update(id string, dish *entities.Dish) (*entities.Dish, error)
	ReduceStock(id string, quantity uint) error
	IncreaseStock(id string, quantity uint) error
	Delete(id, sellerId string) error
}
