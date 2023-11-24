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
	ID         uint   `json:"redeemedId"`
	CouponCode string `json:"couponCode" gorm:"text;references:Coupon;constraint:OnDelete:CASCADE"`
	UserID     uint   `json:"userId"`

	FkUser   User   `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
}
