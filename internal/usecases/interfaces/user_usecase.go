package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IUserUseCase interface {
	SignUp(req *req.UserSignUpReq) (*entities.User, error)
	Login(req *req.UserLoginReq) (*entities.User, error)
	SendOtp(phone string) error
	VerifyOtp(phone string, req *req.UserVerifyOtpReq) error

	SearchDish(search string) (*[]entities.Dish, error)
	GetDishesPage(page, limit string) (*[]entities.Dish, error)
	GetDish(id string) (*entities.Dish, error)

	// TODO: SearchSeller(search string)(*[]entities.Seller, error)
	// TODO: GetSellersPage(page string)(*[]entities.Seller, error)
	// TODO: GetSeller(id string)(*[]entities.Seller, error)

}
