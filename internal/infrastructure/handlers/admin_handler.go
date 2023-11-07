package handlers

import (
	"fmt"

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
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to verify seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully verified seller",
		})

}
func (h *AdminHandler) BlockSeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	if err := h.usecase.BlockSeller(sellerId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to block seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully blocked seller",
		})

}

func (h *AdminHandler) UnblockSeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	if err := h.usecase.UnblockSeller(sellerId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to unblock seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully unblocked seller",
		})

}

func (h *AdminHandler) GetAllUsers(c *fiber.Ctx) error {

	userList, err := h.usecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.UserListRes{
				Status:  "failed",
				Message: "failed to fetch users list",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserListRes{
			Status:   "success",
			Message:  "successfully fetched users' list",
			UserList: *userList,
		})
}

func (h *AdminHandler) BlockUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	if err := h.usecase.BlockUser(userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to block user",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully blocked user",
		})

}

func (h *AdminHandler) UnblockUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	if err := h.usecase.UnblockUser(userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to unblock user",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully unblocked user",
		})

}

func (h *AdminHandler) AddCategory(c *fiber.Ctx) error {
	var req req.CreateCategoryReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to parse request",
				Error:   err.Error(),
			})
	}

	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to parse request",
				Error:   fmt.Sprint(err),
			})
	}

	if err := h.usecase.CreateCategory(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to create category",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully created category",
		})
}

func (h *AdminHandler) EditCategory(c *fiber.Ctx) error {
	categoryId := c.Params("id")
	var req req.UpdateCategoryReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to parse request",
				Error:   err.Error(),
			})
	}

	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to parse request",
				Error:   fmt.Sprint(err),
			})
	}

	if err := h.usecase.UpdateCategory(categoryId, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to update category",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully updated category",
		})

}
