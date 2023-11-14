package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type ViewAllOrdersRes struct {
	Status  string           `json:"status,omitempty"`
	Orders  []entities.Order `json:"orders,omitempty"`
	Message string           `json:"message,omitempty"`
}

type ViewOrderRes struct {
	Status     string               `json:"status,omitempty"`
	Order      entities.Order       `json:"order,omitempty"`
	OrderItems []entities.OrderItem `json:"orderItems,omitempty"`
	Message    string               `json:"message,omitempty"`
}
