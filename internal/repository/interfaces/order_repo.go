package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type IOrderRepository interface {
	CreateOrder(order *entities.Order) error

	FindOrderById(id string) (*entities.Order, error)
	FindOrderItems(id string) (*[]entities.OrderItem, error)

	FindAllOrdersByUserId(userId string) (*[]entities.Order, error)
	FindAllOrdersBySellerId(sellerId string) (*[]entities.Order, error)

	UpdateOrderStatus(id, status string) error

	CancelOrder(id string) error
}
