package usecases

import (
	"errors"
	"fmt"

	reqModels "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository"
	repoInterface "github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type AdminUcase struct {
	adminRepo repoInterface.IAdminRepository
}

type AdminUserInteractor struct {
	AdminRepo repository.AdminRepository
	UserRepo  repository.UserRepository
}

func NewAdminUsecase(repo repoInterface.IAdminRepository) interfaces.IAdminUseCase {
	return &AdminUcase{repo}
}

func (a *AdminUcase) Login(loginReq *reqModels.AdminLoginReq) error {

	admin, err := a.adminRepo.FindByEmail(loginReq.Email)
	if err != nil {
		fmt.Println("DB Error", err.Error())
		return err
	}

	if admin.Password != loginReq.Password {
		fmt.Println("Error is Invalid Password")
		return errors.New("invalid password")
	}

	return nil
}
