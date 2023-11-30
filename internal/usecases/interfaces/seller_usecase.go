package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type ISellerUseCase interface {
	Login(req *req.SellerLoginReq) (string, error)
	SignUp(req *req.SellerSignUpReq) (string, error)

	SearchSeller(search string) (*[]entities.Seller, error)
	GetSellersPage(page, limit string) (*[]entities.Seller, error)
	GetSeller(id string) (*entities.Seller, error)
}
