package usecases

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
)

type adminUserUcase struct {
	AdminRepo interfaces.IAdminRepository
	UserRepo  interfaces.IUserRepository
}

func NewAdminUserUcase(adminRepo interfaces.IAdminRepository, userRepo interfaces.IUserRepository) *adminUserUcase {
	return &adminUserUcase{adminRepo, userRepo}
}

func (repo *adminUserUcase) BlockUser(id string) error {

	if err := repo.UserRepo.Block(id); err != nil {
		return err
	}

	return nil
}
