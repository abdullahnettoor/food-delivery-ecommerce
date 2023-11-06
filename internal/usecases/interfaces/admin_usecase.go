package interfaces

import req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"

type IAdminUseCase interface {
	Login(admin *req.AdminLoginReq) (string, error)
}
