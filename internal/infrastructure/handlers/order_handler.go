package handlers

import (
	"fmt"
	"strings"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type OrderHandler struct {
	usecase interfaces.IOrderUseCase
}

func NewOrderHandler(uc interfaces.IOrderUseCase) *OrderHandler {
	return &OrderHandler{uc}
}

//	@Summary		Place an order
//	@Description	Place a new order for the user
//	@Security		Bearer
//	@Tags			User Order
//	@Accept			json
//	@Produce		json
//	@Param			req	body		req.NewOrderReq	true	"New order request"
//	@Success		200	{object}	res.CommonRes	"Successfully placed order"
//	@Failure		400	{object}	res.CommonRes	"Bad Request"
//	@Failure		401	{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes	"Internal Server Error"
//	@Router			/cart/checkout [post]
func (h *OrderHandler) PlaceOrder(c *fiber.Ctx) error {
	var req req.NewOrderReq

	user := c.Locals("UserModel").(map[string]any)
	id := fmt.Sprint(user["userId"])

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	order, err := h.usecase.PlaceOrder(id, &req)
	if err == e.ErrNotAvailable || err == e.ErrQuantityExceeds {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed. dish is not available / quatity exceeds stock",
			})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to place order",
			})
	}

	if strings.ToUpper(req.PaymentMethod) == "COD" {
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully placed order",
			})
	} else {
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully placed order",
				Result: fiber.Map{
					"key":             viper.GetString("PAYMENT_KEY_ID"),
					"orderId":        order.TransactionID,
					"totalPrice":     order.TotalPrice,
					"deliveryCharge": order.DeliveryCharge,
					"firstName":      fmt.Sprint(user["firstName"]),
					"email":          fmt.Sprint(user["email"]),
					"phone":          fmt.Sprint(user["phone"]),
				},
			})
	}
}

func (h *OrderHandler) PlaceOrderPayOnline(c *fiber.Ctx) error {
	var req req.NewOrderReq

	user := fiber.Map{"userId": "10", "firstName": "Abdullah", "email": "abdullahnettoor@gmail.com", "phone": "9061904860"}
	userId := fmt.Sprint(user["userId"])
	// userId := c.Query("userId")

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	order, err := h.usecase.PlaceOrder(userId, &req)
	if err == e.ErrNotAvailable || err == e.ErrQuantityExceeds {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed. dish is not available / quatity exceeds stock",
			})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to place order",
			})
	}

	if strings.ToLower(req.PaymentMethod) == "online" {
		return c.Status(fiber.StatusOK).
			Render("payment", fiber.Map{
				"ID":             viper.GetString("PAYMENT_KEY_ID"),
				"OrderID":        order.TransactionID,
				"Discount":       order.Discount,
				"TotalPrice":     order.TotalPrice,
				"DeliveryCharge": order.DeliveryCharge,
				"FirstName":      fmt.Sprint(user["firstName"]),
				"Email":          fmt.Sprint(user["email"]),
				"Phone":          fmt.Sprint(user["phone"]),
			})
	} else {
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully placed order",
			})
	}

}

//	@Summary		Verify payment
//	@Description	Verifies a payment using Razorpay details
//	@Security		Bearer
//	@Tags			User Order
//	@Accept			json
//	@Produce		json
//	@Param			req	body		req.VerifyPaymentReq	true	"Payment verification details"
//	@Success		200	{object}	res.CommonRes			"Success: Payment successfully verified"
//	@Failure		401	{object}	res.CommonRes			"Unauthorized Access"
//	@Failure		400	{object}	res.CommonRes			"Bad Request: Failed to verify payment"
//	@Failure		500	{object}	res.CommonRes			"Internal Server Error: Failed to process payment verification"
//	@Router			/order/verifyPayment [post]
func (h *OrderHandler) VerifyPayment(c *fiber.Ctx) error {
	var req req.VerifyPaymentReq

	fmt.Println("Content type is", string(c.Request().Header.ContentType()))
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}

	fmt.Println("Req is", string(c.Body()))

	if err := h.usecase.VerifyPayment(req.OrderID, req.PaymentID, req.RzpSignature); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to verify payment",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully verified payment",
		})
}

//	@Summary		View a specific order
//	@Description	View details of a specific order for the user
//	@Security		Bearer
//	@Tags			User Order
//	@Tags			Seller Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Order ID"
//	@Success		200	{object}	res.ViewOrderRes	"Successfully fetched order"
//	@Failure		400	{object}	res.CommonRes		"Bad Request"
//	@Failure		401	{object}	res.CommonRes		"Unauthorized Access"
//	@Failure		404	{object}	res.CommonRes		"Order not found"
//	@Failure		500	{object}	res.CommonRes		"Internal Server Error"
//	@Router			/orders/{id} [get]
//	@Router			/seller/orders/{id} [get]
func (h *OrderHandler) ViewOrder(c *fiber.Ctx) error {

	id := c.Params("id")

	order, items, err := h.usecase.ViewOrder(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch orders",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.ViewOrderRes{
			Status:     "success",
			Message:    "successfully fetched orders",
			Order:      *order,
			OrderItems: *items,
		})

}

//	@Summary		View all orders for the user
//	@Description	View details of all orders for the user
//	@Security		Bearer
//	@Tags			User Order
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	res.ViewAllOrdersRes	"Successfully fetched orders"
//	@Failure		401	{object}	res.CommonRes			"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes			"Internal Server Error"
//	@Router			/orders [get]
func (h *OrderHandler) ViewOrdersForUser(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	id := fmt.Sprint(user["userId"])

	orders, err := h.usecase.ViewOrdersForUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch orders",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.ViewAllOrdersRes{
			Status:  "success",
			Message: "successfully fetched orders",
			Orders:  *orders,
		})

}

//	@Summary		View all orders for the seller
//	@Description	View details of all orders for the seller
//	@Security		Bearer
//	@Tags			Seller Order
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	res.ViewAllOrdersRes	"Successfully fetched orders"
//	@Failure		401	{object}	res.CommonRes			"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes			"Internal Server Error"
//	@Router			/seller/orders [get]
func (h *OrderHandler) ViewOrdersForSeller(c *fiber.Ctx) error {

	seller := c.Locals("SellerModel").(map[string]any)
	id := fmt.Sprint(seller["sellerId"])

	orders, err := h.usecase.ViewOrdersForSeller(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch orders",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.ViewAllOrdersRes{
			Status:  "success",
			Message: "successfully fetched orders",
			Orders:  *orders,
		})
}

//	@Summary		Update order status
//	@Description	Update the status of a specific order
//	@Security		Bearer
//	@Tags			Seller Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"Order ID"
//	@Param			req	body		req.UpdateOrderStatusReq	true	"Update order status request"
//	@Success		200	{object}	res.CommonRes				"Successfully updated order"
//	@Failure		400	{object}	res.CommonRes				"Bad Request"
//	@Failure		401	{object}	res.CommonRes				"Unauthorized Access"
//	@Failure		404	{object}	res.CommonRes				"Order not found"
//	@Failure		500	{object}	res.CommonRes				"Internal Server Error"
//	@Router			/seller/orders/{id} [patch]
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	var req req.UpdateOrderStatusReq

	id := c.Params("id")

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to parse body",
				Error:   err.Error(),
			})
	}

	if errs := requestvalidation.ValidateRequest(req); errs != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   errs,
				Message: "failed. invalid fields",
			})
	}

	if err := h.usecase.UpdateOrderStatus(id, req.OrderStatus); err != nil {
		if err == e.ErrNotFound {
			return c.Status(fiber.StatusBadRequest).
				JSON(res.CommonRes{
					Status:  "failed",
					Error:   err.Error(),
					Message: "failed to update order",
				})
		}

		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to update order",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully updated order",
		})
}

//	@Summary		Get daily sales report
//	@Description	Fetches the daily sales report for the seller
//	@Security		Bearer
//	@Tags			Seller Sales
//	@Accept			json
//	@Produce		json
//	@Param			filter	query		string			false	"Time intervals"
//	@Success		200		{object}	res.CommonRes	"Success: Daily sales fetched successfully"
//	@Failure		401		{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500		{object}	res.CommonRes	"Internal Server Error: Failed to fetch daily sales"
//	@Router			/seller/sales [get]
func (h *OrderHandler) GetSales(c *fiber.Ctx) error {
	var sales *entities.Sales
	var err error

	seller := c.Locals("SellerModel").(map[string]any)
	id := fmt.Sprint(seller["sellerId"])

	timeIntervals := c.Query("filter")
	switch strings.ToLower(timeIntervals) {
	case "daily":
		sales, err = h.usecase.GetDailySales(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(res.CommonRes{
					Status:  "failed",
					Error:   err.Error(),
					Message: "failed to fetch daily sales",
				})
		}
	case "":
		sales, err = h.usecase.GetTotalSales(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(res.CommonRes{
					Status:  "failed",
					Error:   err.Error(),
					Message: "failed to fetch daily sales",
				})
		}
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully fetched daily sales",
			Result:  *sales,
		})
}
