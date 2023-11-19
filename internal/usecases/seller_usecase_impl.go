package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	hashpassword "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/hash_password"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/spf13/viper"
)

type sellerUsecase struct {
	sellerRepo interfaces.ISellerRepository
	dishRepo   interfaces.IDishRepository
}

func NewSellerUsecase(sellerRepo interfaces.ISellerRepository, dishRepo interfaces.IDishRepository) i.ISellerUseCase {
	return &sellerUsecase{sellerRepo, dishRepo}
}

func (uc *sellerUsecase) Login(req *req.SellerLoginReq) (string, error) {

	seller, err := uc.sellerRepo.FindByEmail(req.Email)
	if err != nil {
		fmt.Println("DB Error", err.Error())
		return "", err
	}

	switch seller.Status {
	case "Pending":
		return "", errors.New("seller is not verified")
	case "Blocked":
		return "", errors.New("seller is blocked")
	case "Rejected":
		return "", errors.New("seller's application is rejected")
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

	_, err := uc.sellerRepo.FindByEmail(req.Email)
	if err != nil && err != e.ErrNotFound {
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

	if err := uc.sellerRepo.Create(&seller); err != nil {
		return "", err
	}

	secret := viper.GetString("KEY")

	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, seller)
	if err != nil {
		return "", err
	}

	return token, nil
}
