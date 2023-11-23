package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type AdminLoginRes struct {
	Status  string `json:"status,omitempty"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type AllCategoriesRes struct {
	Status     string              `json:"status,omitempty"`
	Categories []entities.Category `json:"categories,omitempty"`
	Error      string              `json:"error,omitempty"`
	Message    string              `json:"message,omitempty"`
}

type GetCategoryRes struct {
	Status     string              `json:"status,omitempty"`
	Category entities.Category `json:"category,omitempty"`
	Error      string              `json:"error,omitempty"`
	Message    string              `json:"message,omitempty"`
}
