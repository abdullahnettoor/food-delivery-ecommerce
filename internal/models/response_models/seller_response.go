package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type SellerLoginRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type SellerListRes struct {
	Status     string            `json:"status,omitempty"`
	SellerList []entities.Seller `json:"sellerList,omitempty"`
	Message    string            `json:"message,omitempty"`
}
type SingleSellerRes struct {
	Status  string          `json:"status,omitempty"`
	Seller  entities.Seller `json:"seller,omitempty"`
	Message string          `json:"message,omitempty"`
}
