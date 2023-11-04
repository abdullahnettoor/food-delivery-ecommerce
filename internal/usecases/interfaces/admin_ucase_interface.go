package interfaces

import (
	reqModels "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IAdminUseCase interface {
	Login(admin *reqModels.AdminLoginReq) error
}
