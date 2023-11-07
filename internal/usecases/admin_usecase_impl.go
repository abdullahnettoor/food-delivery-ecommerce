package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/spf13/viper"
)

type adminUcase struct {
	AdminRepo  interfaces.IAdminRepository
	UserRepo   interfaces.IUserRepository
	SellerRepo interfaces.ISellerRepository
}

func NewAdminUsecase(AdminRepo interfaces.IAdminRepository, UserRepo interfaces.IUserRepository, SellerRepo interfaces.ISellerRepository) *adminUcase {
	return &adminUcase{AdminRepo, UserRepo, SellerRepo}
}

func (repo *adminUcase) Login(loginReq *req.AdminLoginReq) (string, error) {

	admin, err := repo.AdminRepo.FindByEmail(loginReq.Email)
	if err != nil {
		fmt.Println("DB Error", err.Error())
		return "", err
	}

	if admin.Password != loginReq.Password {
		fmt.Println("Error is Invalid Password")
		return "", errors.New("invalid password")
	}

	secret := viper.GetString("KEY")
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, admin)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (repo *adminUcase) BlockUser(id string) error {

	if err := repo.UserRepo.Block(id); err != nil {
		return err
	}
	return nil

}

func (repo *adminUcase) UnblockUser(id string) error {
	return repo.UserRepo.Unblock(id)
}

func (repo *adminUcase) GetAllSellers() (*[]entities.Seller, error) {
	return repo.SellerRepo.FindAll()
}

func (repo *adminUcase) VerifySeller(id string) error {
	return repo.SellerRepo.Verify(id)
}

func (repo *adminUcase) BlockSeller(id string) error {
	return repo.SellerRepo.Block(id)
}

func (repo *adminUcase) UnblockSeller(id string) error {
	return repo.SellerRepo.Unblock(id)
}
