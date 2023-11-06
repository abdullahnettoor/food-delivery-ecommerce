package usecases

import (
	"errors"
	"fmt"
	"time"

	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/spf13/viper"
)

type adminUcase struct {
	adminRepo interfaces.IAdminRepository
}

type AdminUserInteractor struct {
	AdminRepo repository.AdminRepository
	UserRepo  repository.UserRepository
}

func NewAdminUsecase(repo interfaces.IAdminRepository) *adminUcase {
	return &adminUcase{repo}
}

func (a *adminUcase) Login(loginReq *req.AdminLoginReq) (string, error) {

	admin, err := a.adminRepo.FindByEmail(loginReq.Email)
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

	// context.WithValue("Admin", claims)

	return token, nil
}
