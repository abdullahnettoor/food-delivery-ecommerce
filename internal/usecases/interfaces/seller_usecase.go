package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type ISellerUseCase interface {
	Login(req *req.SellerLoginReq) (string, error)
	SignUp(req *req.SellerSignUpReq) (string, error)

	GetDish(id, sellerId string) (*entities.Dish, error)
	GetAllDishes(sellerId, category_id string) (*[]entities.Dish, error)
	AddDish(sellerId string, req *req.CreateDishReq) error
	UpdateDish(id, sellerId string, req *req.UpdateDishReq) (*entities.Dish, error)
	DeleteDish(id, sellerId string) error
}
