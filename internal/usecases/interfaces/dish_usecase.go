package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IDishUseCase interface {
	GetDishBySeller(id, sellerId string) (*entities.Dish, error)
	GetAllDishesBySeller(sellerId, categoryId string) (*[]entities.Dish, error)
	AddDish(sellerId string, req *req.CreateDishReq) error
	UpdateDish(id, sellerId string, req *req.UpdateDishReq) (*entities.Dish, error)
	DeleteDish(id, sellerId string) error

	SearchDish(search string) (*[]entities.Dish, error)
	GetDishesPage(categoryId string, page, limit string) (*[]entities.Dish, error)
	GetDish(id string) (*entities.Dish, error)

	ReduceStock(id string, quantity uint) error
}
