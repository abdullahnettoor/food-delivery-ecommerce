package handlers

import (
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	usecase interfaces.IAdminUseCase
}

func NewAdminHandler(uCase interfaces.IAdminUseCase) *AdminHandler {
	return &AdminHandler{uCase}
}

func (h *AdminHandler) Login(c *fiber.Ctx) error {
	var loginReq req.AdminLoginReq

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	if err := requestvalidation.ValidateRequest(loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err,
			})
	}

	token, err := h.usecase.Login(&loginReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"success": "Login Successful",
			"token":   token,
		})
}

func (h *AdminHandler) GetAllSellers(c *fiber.Ctx) error {

	sellerList, err := h.usecase.GetAllSellers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.SellerListRes{
				Status:  "failed",
				Message: "failed to fetch sellers list",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerListRes{
			Status:     "success",
			Message:    "successfully fetched sellers' list",
			SellerList: *sellerList,
		})
}

func (h *AdminHandler) VerifySeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	if err := h.usecase.VerifySeller(sellerId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.AdminCommonRes{
				Status:  "failed",
				Message: "failed to verify seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.AdminCommonRes{
			Status:  "success",
			Message: "successfully verified seller",
		})

}
func (h *AdminHandler) BlockSeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	if err := h.usecase.BlockSeller(sellerId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.AdminCommonRes{
				Status:  "failed",
				Message: "failed to block seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.AdminCommonRes{
			Status:  "success",
			Message: "successfully blocked seller",
		})

}

func (h *AdminHandler) UnblockSeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	if err := h.usecase.UnblockSeller(sellerId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.AdminCommonRes{
				Status:  "failed",
				Message: "failed to unblock seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.AdminCommonRes{
			Status:  "success",
			Message: "successfully unblocked seller",
		})

}
