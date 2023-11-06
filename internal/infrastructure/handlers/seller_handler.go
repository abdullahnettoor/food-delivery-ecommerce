package handlers

import (
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
)

type SellerHandler struct {
	usecase interfaces.ISellerUseCase
}

func NewSellerHandler(uCase interfaces.ISellerUseCase) *SellerHandler {
	return &SellerHandler{uCase}
}

func (h *SellerHandler) SignUp(c *fiber.Ctx) error {
	var signUpReq req.SellerSignUpReq

	if err := c.BodyParser(&signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	if err := requestvalidation.ValidateRequest(signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err,
			})
	}

	token, err := h.usecase.SignUp(&signUpReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerLoginRes{
			Status: "success",
			Token:  token,
		})
}

func (h *SellerHandler) Login(c *fiber.Ctx) error {
	var loginReq req.SellerLoginReq

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
		JSON(res.SellerLoginRes{
			Status: "success",
			Token:  token,
		})
}
