package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IAdminUseCase interface {
	Login(admin *req.AdminLoginReq) (string, error)

	GetAllSellers() (*[]entities.Seller, error)
	VerifySeller(id string) error
	BlockSeller(id string) error
	UnblockSeller(id string) error

	GetAllUsers() (*[]entities.User, error)
	BlockUser(id string) error
	UnblockUser(id string) error
}
