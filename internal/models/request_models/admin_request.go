package req

import "time"

type AdminLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=3"`
}

type CreateCategoryReq struct {
	Name string `json:"name" validate:"required,gte=3"`
}
type UpdateCategoryReq struct {
	Name string `json:"name" validate:"required,gte=3"`
}

type CreateCouponReq struct {
	Code            string    `json:"couponCode" validate:"regexp=^[A-Z]{5}\\d{3}$"`
	Description     string    `json:"description"`
	Type            string    `json:"couponType" validate:"oneof=AMOUNT PERCENTAGE"`
	Discount        uint      `json:"discount" validate:"number"`
	MinimumRequired uint      `json:"minimumAmtRequired" validate:"gte=Discount"`
	MaximumAllowed  uint      `json:"maximumAmtAllowed" validate:"gte=Discount"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	Status          string    `json:"status" validate:"oneof=ACTIVE INACTIVE"`
}
