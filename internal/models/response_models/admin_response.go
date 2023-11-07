package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type AdminLoginRes struct {
	Status  string `json:"status,omitempty"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type AdminCommonRes struct {
	Status  string `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Result  any    `json:"result,omitempty"`
}

type SellerListRes struct {
	Status     string            `json:"status,omitempty"`
	SellerList []entities.Seller `json:"sellerList,omitempty"`
	Error      string            `json:"error,omitempty"`
	Message    string            `json:"message,omitempty"`
}
