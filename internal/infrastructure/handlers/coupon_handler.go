package handlers

import (
	"fmt"

	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	"github.com/gofiber/fiber/v2"
)

type CouponHandler struct {
	uc interfaces.ICouponUseCase
}

func NewCouponHandler(uc interfaces.ICouponUseCase) *CouponHandler {
	return &CouponHandler{uc}
}

// @Summary Get all available coupons for the user
// @Description Fetches all available coupons for the user
// @Security Bearer
// @Tags User Coupons
// @Accept json
// @Produce json
// @Success 200 {object} res.GetAllCouponsForUserRes "Success: Coupons fetched successfully"
// @Success 200 {object} res.CommonRes "Success: No coupons are available"
// @Failure 401 {object} res.CommonRes "Unauthorized Access: User is not authenticated"
// @Failure 500 {object} res.CommonRes "Internal Server Error: Failed to fetch coupons"
// @Router /coupons [get]
func (h *CouponHandler) GetAllCoupons(c *fiber.Ctx) error {

	coupons, err := h.uc.GetCouponsForUser()

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetAllCouponsForUserRes{
				Status:  "success",
				Message: "successfully fetched categories",
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
				Message: "failed to fetch categories",
				Error:   err.Error(),
			})
	}

}

// @Summary Get redeemed coupons by user
// @Description Fetches the coupons redeemed by the user
// @Security Bearer
// @Tags User Coupons
// @Accept json
// @Produce json
// @Success 200 {object} res.GetRedeemedCouponsRes "Success: Redeemed coupons fetched successfully"
// @Success 200 {object} res.CommonRes "Success: No coupons are redeemed"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure 500 {object} res.CommonRes "Internal Server Error: Failed to fetch redeemed coupons"
// @Router /coupons/redeemed [get]
func (h *CouponHandler) GetRedeemedByUser(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)

	userId := fmt.Sprint(user["userId"])

	coupons, err := h.uc.GetRedeemedByUser(userId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.GetRedeemedCouponsRes{
				Status:          "success",
				Message:         "successfully fetched category",
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
				Message: "failed to fetch category",
				Error:   err.Error(),
			})
	}
}
