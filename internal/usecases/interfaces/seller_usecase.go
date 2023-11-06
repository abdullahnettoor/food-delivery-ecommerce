package interfaces

import req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"

type ISellerUseCase interface {
	Login(req *req.SellerLoginReq) (string, error)
	SignUp(req *req.SellerSignUpReq) (string, error)
}
