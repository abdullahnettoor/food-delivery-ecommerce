package usecases

import (
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	hashpassword "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/hash_password"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/spf13/viper"
)

type sellerUsecase struct {
	repo interfaces.ISellerRepository
}

func NewSellerUsecase(repo interfaces.ISellerRepository) *sellerUsecase {
	return &sellerUsecase{repo}
}

func (uc *sellerUsecase) Login(req *req.SellerLoginReq) (string, error) {

	seller, err := uc.repo.FindByEmail(req.Email)
	if err != nil {
		fmt.Println("DB Error", err.Error())
		return "", err
	}

	if err := hashpassword.CompareHashedPassword(seller.Password, req.Password); err != nil {
		fmt.Println("Error is Invalid Password")
		return "", e.ErrInvalidPassword
	}

	secret := viper.GetString("KEY")
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, seller)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *sellerUsecase) SignUp(req *req.SellerSignUpReq) (string, error) {

	_, err := uc.repo.FindByEmail(req.Email)
	if err != nil && err != e.ErrNotFound {
		fmt.Println("UC Seller #1:", err.Error())
		return "", err
	}

	hashedPwd, _ := hashpassword.HashPassword(req.Password)

	seller := entities.Seller{
		Name:        req.Name,
		Description: req.Description,
		Email:       req.Email,
		Password:    hashedPwd,
		PinCode:     req.PinCode,
	}

	if err := uc.repo.Create(&seller); err != nil {
		fmt.Println("UC Seller #2:", err.Error())
		return "", err
	}

	secret := viper.GetString("KEY")

	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, seller)
	if err != nil {
		return "", err
	}

	return token, nil
}
