package handlers

import (
	requestmodels "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	usercase interfaces.IAdminUseCase
}

func NewAdminHandler(uCase interfaces.IAdminUseCase) *AdminHandler {
	return &AdminHandler{uCase}
}

func (a *AdminHandler) Login(c *fiber.Ctx) error {
	var loginReq requestmodels.AdminLoginReq

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := a.usercase.Login(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Login Successfull"})
}
