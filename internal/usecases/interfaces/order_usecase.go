package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IOrderUseCase interface {
	PlaceOrder(userId string, req *req.NewOrderReq) error
	ViewOrder(id string) (*entities.Order, *[]entities.OrderItem, error)
	ViewOrdersForUser(userId string) (*[]entities.Order, error)
	ViewOrdersForSeller(sellerId string) (*[]entities.Order, error)
	UpdateOrderStatus(id, status string) error
	CancelOrder(id string) error
}
