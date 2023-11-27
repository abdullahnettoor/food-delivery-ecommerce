package usecases

import (
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type couponUseCase struct {
	couponRepo interfaces.ICouponRepository
}

func NewCouponUsecase(
	couponRepo interfaces.ICouponRepository,
) i.ICouponUseCase {
	return &couponUseCase{couponRepo}
}

func (uc *couponUseCase) CreateCoupon(req *req.CreateCouponReq) error {
	var coupon = entities.Coupon{
		Code:            req.Code,
		Description:     req.Description,
		Type:            req.Type,
		Discount:        req.Discount,
		MinimumRequired: req.MinimumRequired,
		MaximumAllowed:  req.MaximumAllowed,
		StartDate:       time.Now(),
		Status:          "ACTIVE",
	}

	if c, err := uc.couponRepo.FindByCode(req.Code); err == nil && c.Code == req.Code {
		return e.ErrConflict
	}

	return uc.couponRepo.Create(&coupon)
}

func (uc *couponUseCase) UpdateCouponStatus(id, status string) error {
	switch status {
	case "ACTIVE":
		return uc.couponRepo.UpdateStatus(id, status)
	case "INACTIVE":
		return uc.couponRepo.UpdateStatus(id, status)
	}
	return e.ErrInvalidStatusValue
}

func (uc *couponUseCase) DeleteCoupon(id string) error {
	return uc.couponRepo.Delete(id)
}

func (uc *couponUseCase) GetAllCoupons() (*[]entities.Coupon, error) {
	return uc.couponRepo.FindAll()

}

func (uc *couponUseCase) GetCouponsForUser() (*[]entities.Coupon, error) {
	return uc.couponRepo.FindAllForUser()

}

func (uc *couponUseCase) GetAvailableCouponsForUser(userId string) (*[]entities.Coupon, error) {
	return uc.couponRepo.FindAllAvailableForUser(userId)

}

func (uc *couponUseCase) GetRedeemedByUser(userId string) (*[]entities.RedeemedCoupon, error) {
	return uc.couponRepo.FindRedeemedByUser(userId)
}
