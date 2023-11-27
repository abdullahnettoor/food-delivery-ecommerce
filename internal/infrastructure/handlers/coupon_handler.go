package handlers

import (
	"errors"
	"fmt"

	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
)

type CouponHandler struct {
	uc interfaces.ICouponUseCase
}

func NewCouponHandler(uc interfaces.ICouponUseCase) *CouponHandler {
	return &CouponHandler{uc}
}

// @Summary		Create a new coupon
// @Description	Adds a new coupon to the system
// @Security		Bearer
// @Tags			Admin Coupons
// @Accept			json
// @Produce		json
// @Param			req	body		req.CreateCouponReq	true	"Coupon creation request"
// @Success		200	{object}	res.CommonRes		"Success: Coupon created successfully"
// @Failure		400	{object}	res.CommonRes		"Bad Request: Invalid fields in the request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		409	{object}	res.CommonRes		"Conflict: Coupon with the same code already exists"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error: Failed to create coupon"
// @Router			/admin/coupons/add [post]
func (h *CouponHandler) CreateCoupon(c *fiber.Ctx) error {
	var req req.CreateCouponReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "error occured while parsing body",
				Error:   err.Error(),
			})
	}

	if errs := requestvalidation.ValidateRequest(req); errs != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "invalid fields",
				Error:   fmt.Sprint(errs),
			})
	}

	switch err := h.uc.CreateCoupon(&req); err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetAllCouponsForUserRes{
				Status:  "success",
				Message: "successfully created coupon",
			})

	case e.ErrConflict:
		return c.Status(fiber.StatusConflict).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "coupon already exist",
				Error:   err.Error(),
			})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to create coupon",
				Error:   err.Error(),
			})
	}

}

// @Summary		Update coupon status
// @Description	Updates the status of a coupon in the system
// @Security		Bearer
// @Tags			Admin Coupons
// @Accept			json
// @Produce		json
// @Param			id		path		string			true	"Coupon ID"
// @Param			status	query		string			true	"New status for the coupon"
// @Success		200		{object}	res.CommonRes	"Success: Coupon status updated successfully"
// @Failure		400		{object}	res.CommonRes	"Bad Request: Invalid fields in the request"
// @Failure		401		{object}	res.CommonRes	"Unauthorized Access"
// @Failure		404		{object}	res.CommonRes	"Not Found: Coupon with the given ID not found"
// @Failure		500		{object}	res.CommonRes	"Internal Server Error: Failed to update coupon status"
// @Router			/admin/coupons/{id} [patch]
func (h *CouponHandler) UpdateCouponStatus(c *fiber.Ctx) error {

	couponId := c.Params("id")
	status := c.Query("status")

	if couponId == "" || status == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "invalid fields",
				Error:   errors.New("params or query is invalid").Error(),
			})
	}

	switch err := h.uc.UpdateCouponStatus(couponId, status); err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetAllCouponsForUserRes{
				Status:  "success",
				Message: "successfully updated coupon status",
			})

	case e.ErrNotFound:
		return c.Status(fiber.StatusConflict).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "not found coupon with given id",
				Error:   err.Error(),
			})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to update coupon",
				Error:   err.Error(),
			})
	}

}

// @Summary		Get all coupons
// @Description	Fetches all available coupons
// @Security		Bearer
// @Tags			Admin Coupons
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.GetAllCouponsForUserRes	"Success: Coupons fetched successfully"
// @Success		200	{object}	res.CommonRes				"Success: No coupons are available"
// @Failure		401	{object}	res.CommonRes				"Unauthorized Access: User is not authenticated"
// @Failure		500	{object}	res.CommonRes				"Internal Server Error: Failed to fetch coupons"
// @Router			/admin/coupons [get]
func (h *CouponHandler) GetAllCoupons(c *fiber.Ctx) error {

	coupons, err := h.uc.GetAllCoupons()

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetAllCouponsForUserRes{
				Status:  "success",
				Message: "successfully fetched coupons",
				Coupons: *coupons,
			})

	case e.ErrIsEmpty:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "no coupons are available",
				Error:   e.ErrNotAvailable.Error(),
			})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch coupons",
				Error:   err.Error(),
			})
	}

}

// @Summary		Get all coupons for the user
// @Description	Fetches all available coupons for the user
// @Security		Bearer
// @Tags			User Coupons
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.GetAllCouponsForUserRes	"Success: Coupons fetched successfully"
// @Success		200	{object}	res.CommonRes				"Success: No coupons are available"
// @Failure		401	{object}	res.CommonRes				"Unauthorized Access: User is not authenticated"
// @Failure		500	{object}	res.CommonRes				"Internal Server Error: Failed to fetch coupons"
// @Router			/coupons [get]
func (h *CouponHandler) GetAllCouponsForUser(c *fiber.Ctx) error {

	coupons, err := h.uc.GetCouponsForUser()

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetAllCouponsForUserRes{
				Status:  "success",
				Message: "successfully fetched coupons",
				Coupons: *coupons,
			})

	case e.ErrIsEmpty:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "no coupons are available",
				Error:   e.ErrNotAvailable.Error(),
			})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch coupons",
				Error:   err.Error(),
			})
	}

}

// @Summary		Get available coupons for user
// @Description	Fetches the coupons available for the user
// @Security		Bearer
// @Tags			User Coupons
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.GetAllCouponsForUserRes	"Success: Available coupons fetched successfully"
// @Success		200	{object}	res.CommonRes				"Success: No coupons available"
// @Failure		401	{object}	res.CommonRes				"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes				"Internal Server Error: Failed to fetch available coupons"
// @Router			/coupons/available [get]
func (h *CouponHandler) GetAvailableCouponsForUser(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	coupons, err := h.uc.GetAvailableCouponsForUser(userId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetAllCouponsForUserRes{
				Status:  "success",
				Message: "successfully fetched available coupons",
				Coupons: *coupons,
			})

	case e.ErrIsEmpty:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "no coupons are available",
				Error:   e.ErrNotAvailable.Error(),
			})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch available coupons",
				Error:   err.Error(),
			})
	}
}

// @Summary		Get redeemed coupons by user
// @Description	Fetches the coupons redeemed by the user
// @Security		Bearer
// @Tags			User Coupons
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.GetRedeemedCouponsRes	"Success: Redeemed coupons fetched successfully"
// @Success		200	{object}	res.CommonRes				"Success: No coupons are redeemed"
// @Failure		401	{object}	res.CommonRes				"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes				"Internal Server Error: Failed to fetch redeemed coupons"
// @Router			/coupons/redeemed [get]
func (h *CouponHandler) GetRedeemedByUser(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	coupons, err := h.uc.GetRedeemedByUser(userId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetRedeemedCouponsRes{
				Status:          "success",
				Message:         "successfully fetched coupon",
				RedeemedCoupons: *coupons,
			})

	case e.ErrIsEmpty:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "no coupons are redeemed",
				Error:   err.Error(),
			})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch coupon",
				Error:   err.Error(),
			})
	}
}
