package repository

import (
	"fmt"
	"time"

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
	WHERE user_id = ?
	ORDER BY id DESC
	`,
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
	WHERE seller_id = ?
	AND ((UPPER(payment_status) <> 'PENDING')
	OR (payment_method = 'COD' AND payment_status ILIKE 'PENDING'))
	ORDER BY id DESC
	`,
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

func (repo *orderRepository) UpdateOrderStatus(id, status, paymentStatus string) error {

	if err := repo.DB.Exec(`
	UPDATE orders
	SET status = ?,
	payment_status = ?
	WHERE id = ?`,
		status, paymentStatus, id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *orderRepository) UpdateOrderPaymentStatus(id, status string) error {
	if err := repo.DB.Exec(`
	UPDATE orders
	SET payment_status = ?,
	status = 'Ordered'
	WHERE transaction_id = ?`,
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

func (repo *orderRepository) FindSales(sellerId string, startDate, endDate time.Time) (*entities.Sales, error) {

	var sales entities.Sales

	query := `SELECT
	COUNT(*) as count,
	COALESCE(SUM(total_price), 0) as total_amt 
	FROM orders`
	query += fmt.Sprintf(" WHERE seller_id = %v AND UPPER(status) <> 'PENDING' ", sellerId)

	if !startDate.IsZero() && !endDate.IsZero() {
		start := startDate.Format("2006-01-02 15:04:05-07")
		end := endDate.Format("2006-01-02 15:04:05-07")
		query += " AND order_date BETWEEN '" + start + "' AND '" + end + "' "
	}

	err := repo.DB.Raw(query).Scan(&sales).Error

	if err != nil {
		return nil, err
	}

	return &sales, nil
}
