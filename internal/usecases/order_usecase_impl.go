package usecases

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/services/payment"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	"github.com/google/uuid"
)

type orderUsecase struct {
	cartRepo  interfaces.ICartRepository
	orderRepo interfaces.IOrderRepository
	dishRepo  interfaces.IDishRepository
}

func NewOrderUsecase(cartRepo interfaces.ICartRepository, orderRepo interfaces.IOrderRepository, dishRepo interfaces.IDishRepository) i.IOrderUseCase {
	return &orderUsecase{cartRepo, orderRepo, dishRepo}
}

func (uc *orderUsecase) PlaceOrder(userId string, req *req.NewOrderReq) (*entities.Order, error) {
	var order entities.Order
	var orderItems []entities.OrderItem
	var totalPrice, deliveryCharge, discount float64

	cartItems, err := uc.cartRepo.FindCartItems(userId)
	if err != nil {
		return nil, err
	}

	for _, item := range *cartItems {
		id := fmt.Sprint(item.DishID)
		dish, err := uc.dishRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		if dish.Quantity < item.Quantity {
			return nil, e.ErrQuantityExceeds
		}
		if !dish.Availability {
			return nil, e.ErrNotAvailable
		}
		o := entities.OrderItem{
			DishID:   item.DishID,
			Dish:     *dish,
			Quantity: item.Quantity,
			Price:    dish.Price,
		}
		totalPrice += (dish.Price * float64(item.Quantity))
		orderItems = append(orderItems, o)
	}

	discount = 0
	deliveryCharge = 10
	totalPrice = totalPrice + deliveryCharge - discount
	addressId, _ := strconv.ParseUint(req.AddressID, 10, 0)
	uId, _ := strconv.ParseUint(userId, 10, 0)
	order = entities.Order{
		UserID:         uint(uId),
		AddressID:      uint(addressId),
		SellerID:       orderItems[0].SellerID,
		OrderDate:      time.Now(),
		TransactionID:  uuid.New().String(),
		PaymentMethod:  req.PaymentMethod,
		PaymentStatus: "Pending",
		ItemCount:      uint(len(orderItems)),
		Dishes:         orderItems,
		Discount:       discount,
		DeliveryCharge: deliveryCharge,
		DeliveryDate:   time.Now().Add(time.Minute * 45),
		TotalPrice:     totalPrice,
		Status:         "Ordered",
	}

	if req.PaymentMethod == "Online" {
		rzp := payment.PaymentService{}
		rzpOrder, err := rzp.CreatePaymentOrder(order.TotalPrice)
		if err != nil {
			return nil, err
		}
		order.TransactionID = rzpOrder["id"].(string)
	}

	if err := uc.orderRepo.CreateOrder(&order); err != nil {
		return nil, err
	}

	for _, item := range orderItems {
		id, quantity := fmt.Sprint(item.ID), item.Quantity
		if err := uc.dishRepo.ReduceStock(id, quantity); err != nil {
			return nil, err
		}
	}

	if err := uc.cartRepo.DeleteCart(userId); err != nil {
		return nil, err
	}

	return &order, nil
}

func (uc *orderUsecase) VerifyPayment(orderId, rzpPaymentId, signature string) error {
	rzp := payment.PaymentService{}

	if err := rzp.VerifyPayment(orderId, rzpPaymentId, signature); err != nil {
		return err
	}
	return uc.orderRepo.UpdateOrderPaymentStatus(orderId, "Success")
}

func (uc *orderUsecase) ViewOrder(id string) (*entities.Order, *[]entities.OrderItem, error) {
	var orderItems []entities.OrderItem

	items, err := uc.orderRepo.FindOrderItems(id)
	if err != nil {
		return nil, nil, err
	}

	for _, oItem := range *items {
		id := fmt.Sprint(oItem.DishID)
		dish, err := uc.dishRepo.FindByID(id)
		if err != nil {
			return nil, nil, err
		}
		o := entities.OrderItem{
			DishID:   oItem.DishID,
			Dish:     *dish,
			Quantity: oItem.Quantity,
			Price:    dish.Price,
		}
		orderItems = append(orderItems, o)
	}

	order, err := uc.orderRepo.FindOrderById(id)
	if err != nil {
		return nil, nil, err
	}

	return order, &orderItems, nil
}

func (uc *orderUsecase) ViewOrdersForUser(userId string) (*[]entities.Order, error) {
	return uc.orderRepo.FindAllOrdersByUserId(userId)
}

func (uc *orderUsecase) ViewOrdersForSeller(sellerId string) (*[]entities.Order, error) {
	return uc.orderRepo.FindAllOrdersBySellerId(sellerId)
}

func (uc *orderUsecase) UpdateOrderStatus(id, status string) error {
	return uc.orderRepo.UpdateOrderStatus(id, status)
}

func (uc *orderUsecase) CancelOrder(id string) error {
	return uc.orderRepo.CancelOrder(id)
}
