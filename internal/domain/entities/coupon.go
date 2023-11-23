package entities

import "time"

type Coupon struct {
	ID              uint      `json:"couponId"`
	Code            string    `json:"couponCode" gorm:"unique"`
	Description     string    `json:"description"`
	Type            string    `json:"couponType" gorm:"default:AMOUNT"`
	Discount        uint      `json:"discount"`
	MinimumRequired uint      `json:"minimumAmtRequired"`
	MaximumAllowed  uint      `json:"maximumAmtAllowed"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	Status          string    `json:"status" gorm:"default:ACTIVE"`
}

type RedeemedCoupon struct {
	ID         uint   `json:"-"`
	CouponCode string `json:"-"`
	UserID     uint   `json:"-"`

	FkUser   User   `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	FkCoupon Coupon `json:"-" gorm:"foreignkey:CouponCode;constraint:OnDelete:CASCADE"`
}
