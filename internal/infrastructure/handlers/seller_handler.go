package handlers

import (
	"fmt"

	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
)

type SellerHandler struct {
	usecase   interfaces.ISellerUseCase
	dishUcase interfaces.IDishUseCase
}

func NewSellerHandler(uCase interfaces.ISellerUseCase, dishUcase interfaces.IDishUseCase) *SellerHandler {
	return &SellerHandler{uCase, dishUcase}
}

// @Summary		Seller Sign Up
// @Description	Register a new seller
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			req	body		req.SellerSignUpReq	true	"Seller sign-up request"
// @Success		200	{object}	res.SellerLoginRes	"Successfully signed up"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/seller/register [post]
func (h *SellerHandler) SignUp(c *fiber.Ctx) error {
	var signUpReq req.SellerSignUpReq

	if err := c.BodyParser(&signUpReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if err := requestvalidation.ValidateRequest(signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	token, err := h.usecase.SignUp(&signUpReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to register",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerLoginRes{
			Status:  "success",
			Message: "successfully signed up",
			Token:   token,
		})
}

// @Summary		Seller Login
// @Description	Authenticate and log in as a seller
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			req	body		req.SellerLoginReq	true	"Seller login request"
// @Success		200	{object}	res.SellerLoginRes	"Successfully logged in"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/seller/login [post]
func (h *SellerHandler) Login(c *fiber.Ctx) error {
	var loginReq req.SellerLoginReq

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to parse body",
				Error:   err.Error(),
			})
	}
	if err := requestvalidation.ValidateRequest(loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	token, err := h.usecase.Login(&loginReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "failed to Login",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerLoginRes{
			Status:  "success",
			Message: "successfully logged in",
			Token:   token,
		})
}
