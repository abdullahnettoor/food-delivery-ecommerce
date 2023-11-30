package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IUserUseCase interface {
	SignUp(req *req.UserSignUpReq) (*string, error)
	Login(req *req.UserLoginReq) (*string, error)
	SendOtp(phone string) error
	VerifyOtp(phone string, req *req.UserVerifyOtpReq) error
	ForgotPassword(req *req.ForgotPasswordReq) error 
	ResetPassword(req *req.ResetPasswordReq) error

	GetUserDetails(id string) (*entities.User, error)
	UpdateUserDetails(id string, req *req.UpdateUserDetailsReq) (*entities.User, error)
	ChangePassword(id string, req *req.ChangePasswordReq) error

	AddAddress(id string, req *req.NewAddressReq) error
	UpdateAddress(userId, addressId string, req *req.UpdateAddressReq) error 
	ViewAddress(id, userId string) (*entities.Address, error)
	ViewAllAddresses(userId string) (*[]entities.Address, error)

}
