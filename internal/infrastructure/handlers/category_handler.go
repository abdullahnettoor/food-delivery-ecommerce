package handlers

import (
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	uc interfaces.ICategoryUseCase
}

func NewCategoryHandler(uCase interfaces.ICategoryUseCase) *CategoryHandler {
	return &CategoryHandler{uCase}
}

// @Summary		Get all categories
// @Description	Retrieve a list of all categories
// @Tags			Category
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.AllCategoriesRes	"Successful operation"
// @Failure		500	{object}	res.CommonRes			"Internal Server Error"
// @Router			/categories [get]
func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {

	categories, err := h.uc.GetAllCategory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch categories",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.AllCategoriesRes{
			Status:     "success",
			Message:    "successfully fetched categories",
			Categories: *categories,
		})
}

// @Summary		Get all categories
// @Description	Retrieve a list of all categories
// @Tags			Category
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Category ID"	int
// @Success		200	{object}	res.AllCategoriesRes	"Successful operation"
// @Failure		500	{object}	res.CommonRes			"Internal Server Error"
// @Router			/categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	categoryId := c.Params("id")

	category, err := h.uc.GetCategory(categoryId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch category",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.GetCategoryRes{
			Status:   "success",
			Message:  "successfully fetched category",
			Category: *category,
		})
}
