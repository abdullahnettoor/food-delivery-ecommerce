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

// @Summary	Admin login
// @Tags		Admin
// @Accept		json
// @Produce	json
// @Param		adminLoginReq	body		req.AdminLoginReq	true	"Admin Login Request"
// @Success	200				{object}	res.AdminLoginRes	"Successful login"
// @Failure	400				{object}	res.CommonRes		"Bad Request"
// @Failure	500				{object}	res.CommonRes		"Internal Server Error"
// @Router		/admin/login [post]
func (h *AdminHandler) Login(c *fiber.Ctx) error {
	var loginReq req.AdminLoginReq

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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
				Message: "failed to validate body",
				Error:   fmt.Sprint(err),
			})
	}

	token, err := h.usecase.Login(&loginReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to login",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.AdminLoginRes{
			Status:  "success",
			Token:   token,
			Message: "successfully logged in",
		})
}

// @Summary		Get all sellers
// @Description	Get a list of all sellers
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.SellerListRes	"Successful operation"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Router			/admin/sellers [get]
func (h *AdminHandler) GetAllSellers(c *fiber.Ctx) error {

	sellerList, err := h.usecase.GetAllSellers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch sellers list",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerListRes{
			Status:     "success",
			Message:    "successfully fetched sellers' list",
			SellerList: *sellerList,
		})
}

// @Summary		Verify a seller
// @Description	Verify a specific seller by ID
// @Security		Bearer
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Seller ID"	int
// @Success		200	{object}	res.CommonRes	"Seller successfully verified"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Router			/admin/sellers/{id}/verify [patch]
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

// @Summary		Block a seller
// @Description	Block a specific seller by ID
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Seller ID"	int
// @Success		200	{object}	res.CommonRes	"Seller successfully blocked"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/admin/sellers/{id}/block [patch]
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

// @Summary		Unblock a seller
// @Description	Unblock a specific seller by ID
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Seller ID"	int
// @Success		200	{object}	res.CommonRes	"Seller successfully unblocked"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/admin/sellers/{id}/unblock [patch]
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

// @Summary		Get all users
// @Description	Get a list of all users
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.UserListRes	"Successful operation"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/admin/users [get]
func (h *AdminHandler) GetAllUsers(c *fiber.Ctx) error {

	userList, err := h.usecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
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

// @Summary		Block a user
// @Description	Block a specific user by ID
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"User ID"	int
// @Success		200	{object}	res.CommonRes	"User successfully blocked"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/admin/users/{id}/block [patch]
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

// @Summary		Unblock a user
// @Description	Unblock a specific user by ID
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"User ID"	int
// @Success		200	{object}	res.CommonRes	"User successfully unblocked"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/admin/users/{id}/unblock [patch]
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

// @Summary		Add a category
// @Description	Create a new category
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			req	body		req.CreateCategoryReq	true	"Category creation request"
// @Success		200	{object}	res.CommonRes			"Category successfully created"
// @Failure		400	{object}	res.CommonRes			"Bad Request"
// @Failure		401	{object}	res.CommonRes			"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes			"Internal Server Error"
// @Router			/admin/categories/addCategory [post]
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

// @Summary		Edit a category
// @Description	Update an existing category by ID
// @Security		Bearer
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path		string					true	"Category ID"	int
// @Param			req	body		req.UpdateCategoryReq	true	"Category update request"
// @Success		200	{object}	res.CommonRes			"Category successfully updated"
// @Failure		400	{object}	res.CommonRes			"Bad Request"
// @Failure		401	{object}	res.CommonRes			"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes			"Internal Server Error"
// @Router			/admin/categories/{id}/edit [patch]
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

