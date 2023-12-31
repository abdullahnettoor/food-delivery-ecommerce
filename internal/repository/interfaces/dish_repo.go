package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type IDishRepository interface {
	Search(search, sellerId string) (*[]entities.Dish, error)
	FindPageWise(sellerId, categoryId string, page, limit uint) (*[]entities.Dish, error)
	FindByID(id string) (*entities.Dish, error)
	FindBySeller(sellerId, category_id string) (*[]entities.Dish, error)
	FindBySellerAndID(id, sellerId string) (*entities.Dish, error)
	Create(dish *entities.Dish) error
	Update(id string, dish *entities.Dish) (*entities.Dish, error)
	ReduceStock(id string, quantity uint) error
	IncreaseStock(id string, quantity uint) error
	Delete(id, sellerId string) error
}
