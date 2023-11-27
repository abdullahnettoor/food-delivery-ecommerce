package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type ICouponUseCase interface {
	CreateCoupon(req *req.CreateCouponReq) error
	UpdateCouponStatus(id, status string) error
	DeleteCoupon(id string) error
	GetAllCoupons() (*[]entities.Coupon, error)
	
	GetCouponsForUser() (*[]entities.Coupon, error)
	GetAvailableCouponsForUser(userId string) (*[]entities.Coupon, error)
	GetRedeemedByUser(userId string) (*[]entities.RedeemedCoupon, error)
}
