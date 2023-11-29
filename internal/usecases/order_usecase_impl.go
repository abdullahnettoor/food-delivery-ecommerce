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
	cartRepo   interfaces.ICartRepository
	orderRepo  interfaces.IOrderRepository
	dishUcase  i.IDishUseCase
	couponRepo interfaces.ICouponRepository
}

func NewOrderUsecase(
	cartRepo interfaces.ICartRepository,
	orderRepo interfaces.IOrderRepository,
	dishUcase i.IDishUseCase,
	couponRepo interfaces.ICouponRepository,
) i.IOrderUseCase {
	return &orderUsecase{cartRepo, orderRepo, dishUcase, couponRepo}
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
		dish, err := uc.dishUcase.GetDish(id)
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
			DishID:    item.DishID,
			Dish:      *dish,
			Quantity:  item.Quantity,
			SalePrice: dish.SalePrice,
		}
		totalPrice += (dish.SalePrice * float64(item.Quantity))
		orderItems = append(orderItems, o)
	}

	if req.CouponCode != "" {
		_, err := uc.couponRepo.FindRedeemed(userId, req.CouponCode)
		if err != nil && err != e.ErrNotFound {
			return nil, err
		}
		if err == e.ErrNotFound {
			coupon, err := uc.couponRepo.FindByCode(req.CouponCode)
			if err != nil {
				return nil, err
			}
			if coupon.Status != "ACTIVE" {
				return nil, e.ErrInvalidCoupon
			}
			if coupon.MinimumRequired > uint(totalPrice) {
				return nil, e.ErrCouponNotApplicable
			}
			discount = float64(coupon.Discount)
			if discount >= float64(coupon.MaximumAllowed) {
				discount = float64(coupon.MaximumAllowed)
			}
		} else {
			return nil, e.ErrCouponAlreadyRedeemed
		}
	}

	if totalPrice < 500 {
		deliveryCharge = totalPrice * .1
	}

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
		PaymentStatus:  "Pending",
		ItemCount:      uint(len(orderItems)),
		Dishes:         orderItems,
		Discount:       discount,
		DeliveryCharge: deliveryCharge,
		DeliveryDate:   time.Now().Add(time.Minute * 45),
		TotalPrice:     totalPrice,
		Status:         "Ordered",
	}
	if req.CouponCode != "" {
		order.CouponCode = req.CouponCode
	}

	if req.PaymentMethod == "ONLINE" {
		rzp := payment.PaymentService{}
		rzpOrder, err := rzp.CreatePaymentOrder(order.TotalPrice)
		if err != nil {
			return nil, err
		}
		order.Status = "Pending"
		order.TransactionID = rzpOrder["id"].(string)
	}

	if err := uc.orderRepo.CreateOrder(&order); err != nil {
		return nil, err
	}

	if req.CouponCode != "" {
		if err := uc.couponRepo.CreateRedeemed(userId, req.CouponCode); err != nil {
			return nil, err
		}
	}

	for _, item := range orderItems {
		id, quantity := fmt.Sprint(item.DishID), item.Quantity
		if err := uc.dishUcase.ReduceStock(id, quantity); err != nil {
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
		dish, err := uc.dishUcase.GetDish(id)
		if err != nil {
			return nil, nil, err
		}
		o := entities.OrderItem{
			DishID:    oItem.DishID,
			Dish:      *dish,
			Quantity:  oItem.Quantity,
			SalePrice: oItem.SalePrice,
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

func (uc *orderUsecase) GetDailySalesReport(sellerId string) (*entities.Sales, error) {

	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := startDate.Add(time.Hour * 24)

	return uc.orderRepo.FindSales(sellerId, startDate, endDate)
}

func (uc *orderUsecase) GetSalesReportByRange(sellerId string, startDate, endDate time.Time) (*entities.Sales, error) {

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	return uc.orderRepo.FindSales(sellerId,startDate, endDate)
}
