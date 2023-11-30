package usecases

import (
	"errors"
	"fmt"
	"strconv"
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

type sellerUcase struct {
	repo interfaces.ISellerRepository
}

func NewSellerUsecase(repo interfaces.ISellerRepository) i.ISellerUseCase {
	return &sellerUcase{repo}
}

func (uc *sellerUcase) Login(req *req.SellerLoginReq) (string, error) {

	seller, err := uc.repo.FindByEmail(req.Email)
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
	token, _, err := jwttoken.CreateToken(secret, "seller", time.Hour*24, seller)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *sellerUcase) SignUp(req *req.SellerSignUpReq) (string, error) {

	_, err := uc.repo.FindByEmail(req.Email)
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

	if err := uc.repo.Create(&seller); err != nil {
		return "", err
	}

	secret := viper.GetString("KEY")

	token, _, err := jwttoken.CreateToken(secret, "seller", time.Hour*24, seller)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *sellerUcase) SearchSeller(search string) (*[]entities.Seller, error) {
	return uc.repo.SearchVerified(search)
}

func (uc *sellerUcase) GetSellersPage(page, limit string) (*[]entities.Seller, error) {
	p, err := strconv.ParseUint(page, 10, 0)
	if err != nil {
		return nil, err
	}
	l, err := strconv.ParseUint(limit, 10, 0)
	if err != nil {
		return nil, err
	}

	return uc.repo.FindPageWise(uint(p), uint(l))
}

func (uc *sellerUcase) GetSeller(id string) (*entities.Seller, error) {
	return uc.repo.FindVerifiedByID(id)
}
