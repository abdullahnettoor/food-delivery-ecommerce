package repository

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.IOrderRepository {
	return &orderRepository{db}
}

func (repo *orderRepository) CreateOrder(order *entities.Order) error {

	if err := repo.DB.Create(&order).Error; err != nil {
		return err
	}

	var items []entities.OrderItem
	for _, item := range order.Dishes {
		item.OrderID = order.ID
		items = append(items, item)
		
	}

	if err := repo.DB.Create(&items).Error; err != nil {
		return err
	}

	return nil
}

func (repo *orderRepository) FindOrderById(id string) (*entities.Order, error) {
	var order entities.Order

	res := repo.DB.Raw(`
	SELECT *
	FROM orders
	WHERE id = ?`,
		id).Scan(&order)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &order, nil
}

func (repo *orderRepository) FindAllOrdersByUserId(userId string) (*[]entities.Order, error) {
	var orderList []entities.Order

	res := repo.DB.Raw(`
	SELECT *
	FROM orders
	WHERE user_id = ?`,
		userId).Scan(&orderList)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &orderList, nil
}

func (repo *orderRepository) FindAllOrdersBySellerId(sellerId string) (*[]entities.Order, error) {
	var orderList []entities.Order

	res := repo.DB.Raw(`
	SELECT *
	FROM orders
	WHERE seller_id = ?`,
		sellerId).Scan(&orderList)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &orderList, nil
}

func (repo *orderRepository) FindOrderItems(id string) (*[]entities.OrderItem, error) {
	var orderItems []entities.OrderItem

	res := repo.DB.Raw(`
	SELECT *
	FROM order_items
	WHERE order_id = ?`,
		id).Scan(&orderItems)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &orderItems, nil
}

func (repo *orderRepository) UpdateOrderStatus(id, status string) error {
	if err := repo.DB.Exec(`
	UPDATE orders
	SET status = ?
	WHERE id = ?`,
		status, id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *orderRepository) CancelOrder(id string) error {
	if err := repo.DB.Exec(`
	UPDATE orders
	SET status = 'Cancelled'
	WHERE id = ?`,
		id).Error; err != nil {
		return err
	}

	return nil
}
