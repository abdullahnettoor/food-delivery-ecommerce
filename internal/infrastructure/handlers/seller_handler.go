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
	usecase interfaces.ISellerUseCase
}

func NewSellerHandler(uCase interfaces.ISellerUseCase) *SellerHandler {
	return &SellerHandler{uCase}
}

func (h *SellerHandler) SignUp(c *fiber.Ctx) error {
	var signUpReq req.SellerSignUpReq

	if err := c.BodyParser(&signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
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
			Status: "success",
			Token:  token,
		})
}

func (h *SellerHandler) CreateDish(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)
	var req req.CreateDishReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse request",
			})
	}
	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "invalid inputs",
			})
	}

	if err := h.usecase.AddDish(fmt.Sprint(seller["sellerId"]), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to add new dish",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully created new dish",
		})
}

func (h *SellerHandler) UpdateDish(c *fiber.Ctx) error {
	dishId := c.Params("id")
	seller := c.Locals("SellerModel").(map[string]any)
	var req req.UpdateDishReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse request",
			})
	}
	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "invalid inputs",
			})
	}

	dish, err := h.usecase.UpdateDish(dishId, fmt.Sprint(seller["sellerId"]), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to add new dish",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully created new dish",
			Result:  dish,
		})
}

func (h *SellerHandler) GetDish(c *fiber.Ctx) error {
	dishId := c.Params("id")
	seller := c.Locals("SellerModel").(map[string]any)

	dish, err := h.usecase.GetDish(dishId, fmt.Sprint(seller["sellerId"]))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch dish",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SingleDishRes{
			Status:  "success",
			Message: "successfully fetched dish",
			Dish:    *dish,
		})
}

func (h *SellerHandler) GetAllDish(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)

	dishList, err := h.usecase.GetAllDishes(fmt.Sprint(seller["sellerId"]))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch dish",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.DishListRes{
			Status:   "success",
			Message:  "successfully fetched dishes",
			DishList: *dishList,
		})
}

func (h *SellerHandler) DeleteDish(c *fiber.Ctx) error {
	dishId := c.Params("id")
	seller := c.Locals("SellerModel").(map[string]any)

	err := h.usecase.DeleteDish(dishId, fmt.Sprint(seller["sellerId"]))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to delete dish",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully deleted dish",
		})
}
