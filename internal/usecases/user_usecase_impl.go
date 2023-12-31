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
	otphelper "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/twilio"
	"github.com/spf13/viper"
)

type userUcase struct {
	userRepo   interfaces.IUserRepository
}

func NewUserUsecase(userRepo interfaces.IUserRepository) i.IUserUseCase {
	return &userUcase{userRepo}
}

func (uc *userUcase) SignUp(req *req.UserSignUpReq) (*string, error) {

	_, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil && err != e.ErrNotFound {
		return nil, err
	}
	if err != e.ErrNotFound {
		return nil, e.ErrConflict
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

	secret := viper.GetString("KEY")
	fmt.Println("Key is", secret)
	token, _, err := jwttoken.CreateToken(secret, "user", time.Hour*24, *newUser)
	if err != nil {
		return nil, err
	}

	return &token, nil
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

func (uc *userUcase) Login(req *req.UserLoginReq) (*string, error) {

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

	secret := viper.GetString("KEY")
	fmt.Println("Key is", secret)
	token, _, err := jwttoken.CreateToken(secret, "user", time.Hour*24, *user)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (uc *userUcase) GetUserDetails(id string) (*entities.User, error) {
	return uc.userRepo.FindByID(id)
}

func (uc *userUcase) UpdateUserDetails(id string, req *req.UpdateUserDetailsReq) (*entities.User, error) {
	_, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil && err != e.ErrNotFound {
		return nil, err
	}

	user := entities.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	return uc.userRepo.Update(id, &user)
}

func (uc *userUcase) ChangePassword(id string, req *req.ChangePasswordReq) error {

	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	if err := hashpassword.CompareHashedPassword(user.Password, req.Password); err != nil {
		return err
	}

	newPassword, err := hashpassword.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return uc.userRepo.ChangePassword(id, newPassword)
}

func (uc *userUcase) ForgotPassword(req *req.ForgotPasswordReq) error {
	user, err := uc.userRepo.FindByPhone(req.Phone)
	if err != nil {
		return err
	}

	return otphelper.SendOtp(user.Phone)
}

func (uc *userUcase) ResetPassword(req *req.ResetPasswordReq) error {
	if ok, err := otphelper.VerifyOtp(req.Phone, req.Otp); !ok || err != nil {
		return err
	}

	user, err := uc.userRepo.FindByPhone(req.Phone)
	if err != nil {
		return err
	}

	hashPwd, err := hashpassword.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return uc.userRepo.ChangePassword(fmt.Sprint(user.ID), hashPwd)
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
