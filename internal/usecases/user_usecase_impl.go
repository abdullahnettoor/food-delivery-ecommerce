package usecases

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	hashpassword "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/hash_password"
	otphelper "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/twilio"
)

type userUcase struct {
	userRepo   interfaces.IUserRepository
	sellerRepo interfaces.ISellerRepository
}

func NewUserUsecase(userRepo interfaces.IUserRepository, sellerRepo interfaces.ISellerRepository) i.IUserUseCase {
	return &userUcase{userRepo, sellerRepo}
}

func (uc *userUcase) SignUp(req *req.UserSignUpReq) (*entities.User, error) {

	_, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil && err != e.ErrNotFound {
		return nil, err
	}

	hashedPwd, err := hashpassword.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	var user = entities.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPwd,
	}

	newUser, err := uc.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	err = otphelper.SendOtp(user.Phone)
	if err != nil {
		uc.userRepo.DeleteByPhone(user.Phone)
		return nil, err
	}

	return newUser, nil
}

func (uc *userUcase) VerifyOtp(phone string, req *req.UserVerifyOtpReq) error {

	if ok, err := otphelper.VerifyOtp(phone, req.Otp); err != nil || !ok {
		fmt.Println("Inside otp helper")
		return errors.New("invalid otp")
	}
	if err := uc.userRepo.Verify(phone); err != nil {
		fmt.Println("Inside verify user")
		return err
	}

	return nil
}

func (uc *userUcase) SendOtp(phone string) error {
	if err := otphelper.SendOtp(phone); err != nil {
		return err
	}
	return nil
}

func (uc *userUcase) Login(req *req.UserLoginReq) (*entities.User, error) {

	user, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	switch user.Status {
	case "Blocked":
		return nil, errors.New("user is blocked")
	case "Pending":
		return nil, errors.New("user need to verify otp")
	case "Deleted":
		return nil, e.ErrNotFound
	}

	if err := hashpassword.CompareHashedPassword(user.Password, req.Password); err != nil {
		return nil, e.ErrInvalidPassword
	}

	return user, nil
}

func (uc *userUcase) AddAddress(id string, req *req.NewAddressReq) error {

	userId, _ := strconv.ParseUint(id, 10, 0)

	address := entities.Address{
		UserID:    uint(userId),
		Name:      req.Name,
		HouseName: req.HouseName,
		Street:    req.Street,
		District:  req.District,
		State:     req.State,
		PinCode:   req.PinCode,
		Phone:     req.Phone,
	}

	return uc.userRepo.AddAddress(&address)
}

func (uc *userUcase) UpdateAddress(userId, addressId string, req *req.UpdateAddressReq) error {

	uId, _ := strconv.ParseUint(userId, 10, 0)
	aId, _ := strconv.ParseUint(addressId, 10, 0)

	address := entities.Address{
		ID:        uint(aId),
		UserID:    uint(uId),
		Name:      req.Name,
		HouseName: req.HouseName,
		Street:    req.Street,
		District:  req.District,
		State:     req.State,
		PinCode:   req.PinCode,
		Phone:     req.Phone,
	}

	return uc.userRepo.UpdateAddress(addressId, &address)
}

func (uc *userUcase) ViewAddress(id, userId string) (*entities.Address, error) {
	return uc.userRepo.FindAddressByUserID(id, userId)
}

func (uc *userUcase) ViewAllAddresses(userId string) (*[]entities.Address, error) {
	return uc.userRepo.FindAllAddressByUserID(userId)
}

func (uc *userUcase) SearchSeller(search string) (*[]entities.Seller, error) {
	return uc.sellerRepo.SearchVerified(search)
}

func (uc *userUcase) GetSellersPage(page, limit string) (*[]entities.Seller, error) {
	p, err := strconv.ParseUint(page, 10, 0)
	if err != nil {
		return nil, err
	}
	l, err := strconv.ParseUint(limit, 10, 0)
	if err != nil {
		return nil, err
	}

	return uc.sellerRepo.FindPageWise(uint(p), uint(l))
}

func (uc *userUcase) GetSeller(id string) (*entities.Seller, error) {
	return uc.sellerRepo.FindVerifiedByID(id)
}
