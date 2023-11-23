package usecases

import (
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/spf13/viper"
)

type adminUcase struct {
	AdminRepo    interfaces.IAdminRepository
	UserRepo     interfaces.IUserRepository
	SellerRepo   interfaces.ISellerRepository
	CategoryRepo interfaces.ICategoryRepository
}

func NewAdminUsecase(
	AdminRepo interfaces.IAdminRepository,
	UserRepo interfaces.IUserRepository,
	SellerRepo interfaces.ISellerRepository,
	CategoryRepo interfaces.ICategoryRepository) *adminUcase {
	return &adminUcase{AdminRepo, UserRepo, SellerRepo, CategoryRepo}
}

func (uc *adminUcase) Login(loginReq *req.AdminLoginReq) (string, error) {

	admin, err := uc.AdminRepo.FindByEmail(loginReq.Email)
	if err != nil {
		fmt.Println("DB Error", err.Error())
		return "", err
	}

	if admin.Password != loginReq.Password {
		fmt.Println("Error is Invalid Password")
		return "", e.ErrInvalidPassword
	}

	secret := viper.GetString("KEY")
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, admin)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *adminUcase) GetAllSellers() (*[]entities.Seller, error) {
	return uc.SellerRepo.FindAll()
}

func (uc *adminUcase) VerifySeller(id string) error {
	return uc.SellerRepo.Verify(id)
}

func (uc *adminUcase) BlockSeller(id string) error {
	return uc.SellerRepo.Block(id)
}

func (uc *adminUcase) UnblockSeller(id string) error {
	return uc.SellerRepo.Unblock(id)
}

func (uc *adminUcase) GetAllUsers() (*[]entities.User, error) {
	return uc.UserRepo.FindAll()
}

func (uc *adminUcase) BlockUser(id string) error {

	if err := uc.UserRepo.Block(id); err != nil {
		return err
	}
	return nil

}

func (uc *adminUcase) UnblockUser(id string) error {
	return uc.UserRepo.Unblock(id)
}

func (uc *adminUcase) CreateCategory(req *req.CreateCategoryReq) error {
	category := entities.Category{
		Name: req.Name,
	}

	return uc.CategoryRepo.Create(&category)
}

func (uc *adminUcase) UpdateCategory(id string, req *req.UpdateCategoryReq) error {
	return uc.CategoryRepo.Update(id, req.Name)
}

