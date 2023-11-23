package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type GetAllCouponsForUserRes struct {
	Status  string            `json:"status,omitempty"`
	Coupons []entities.Coupon `json:"coupons,omitempty"`
	Error   string            `json:"error,omitempty"`
	Message string            `json:"message,omitempty"`
}

type GetRedeemedCouponsRes struct {
	Status          string                     `json:"status,omitempty"`
	RedeemedCoupons []entities.RedeemedCoupon `json:"redeemedCoupons,omitempty"`
	Error           string                     `json:"error,omitempty"`
	Message         string                     `json:"message,omitempty"`
}
