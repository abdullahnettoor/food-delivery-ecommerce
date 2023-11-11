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
	hashpassword "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/hash_password"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/spf13/viper"
)

type sellerUsecase struct {
	sellerRepo interfaces.ISellerRepository
	dishRepo   interfaces.IDishRepository
}

func NewSellerUsecase(sellerRepo interfaces.ISellerRepository, dishRepo interfaces.IDishRepository) *sellerUsecase {
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

	if err := uc.sellerRepo.Create(&seller); err != nil {
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

func (uc *sellerUsecase) AddDish(sellerId string, req *req.CreateDishReq) error {

	id, err := strconv.ParseInt(sellerId, 10, 32)
	if err != nil {
		return err
	}

	newDish := entities.Dish{
		SellerID:     uint(id),
		Name:         req.Name,
		Description:  req.Description,
		ImageUrl:     req.ImageUrl,
		Price:        req.Price,
		Quantity:     req.Quantity,
		CategoryID:   req.CategoryID,
		IsVeg:        req.IsVeg,
		Availability: req.Availability,
	}
	return uc.dishRepo.Create(&newDish)
}

func (uc *sellerUsecase) UpdateDish(dishId, sellerId string, req *req.UpdateDishReq) (*entities.Dish, error) {

	id, err := strconv.ParseInt(sellerId, 10, 32)
	if err != nil {
		return nil, err
	}

	updatedDish := entities.Dish{
		SellerID:     uint(id),
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		Quantity:     req.Quantity,
		CategoryID:   req.CategoryID,
		IsVeg:        req.IsVeg,
		Availability: req.Availability,
	}
	return uc.dishRepo.Update(dishId, &updatedDish)
}

func (uc *sellerUsecase) GetAllDishes(sellerId string) (*[]entities.Dish, error) {
	return uc.dishRepo.FindBySeller(sellerId)
}

func (uc *sellerUsecase) GetDish(id, sellerId string) (*entities.Dish, error) {
	return uc.dishRepo.FindBySellerAndID(id, sellerId)
}

func (uc *sellerUsecase) DeleteDish(id, sellerId string) error {
	return uc.dishRepo.Delete(id, sellerId)
}
