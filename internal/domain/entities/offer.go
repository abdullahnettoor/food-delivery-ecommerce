package entities

import "time"

type CategoryOffer struct {
	ID         uint      `json:"offerId"`
	Title      string    `json:"offerTitle"`
	SellerID   uint      `json:"sellerId"`
	CategoryID uint      `json:"categoryId"`
	Percentage uint      `json:"offerPercentage"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Status     string    `json:"status"`
}
