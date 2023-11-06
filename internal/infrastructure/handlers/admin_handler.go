package handlers

import (
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
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
		return c.Status(fiber.StatusBadRequest).
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
		return c.Status(fiber.StatusBadRequest).
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
