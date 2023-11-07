package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type ISellerUseCase interface {
	Login(req *req.SellerLoginReq) (string, error)
	SignUp(req *req.SellerSignUpReq) (string, error)

	AddDish(restaurantId string, req *req.CreateDishReq) error
	UpdateDish(id, restaurantId string, req *req.UpdateDishReq) (*entities.Dish, error)
}
