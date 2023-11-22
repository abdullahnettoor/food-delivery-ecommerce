package handlers

import (
	"fmt"

	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	"github.com/gofiber/fiber/v2"
)

type FavHandler struct {
	uc interfaces.IFavouriteUseCase
}

func NewFavHandler(uc interfaces.IFavouriteUseCase) *FavHandler {
	return &FavHandler{uc}
}

// @Summary Add item to favorites
// @Description Adds a dish to the user's favorites
// @Security Bearer
// @Tags User Favorites
// @Accept json
// @Produce json
// @Param id path string true "Dish ID" 
// @Success 200 {object} res.CommonRes "Success: Dish added to favorites successfully"
// @Failure 400 {object} res.CommonRes "Bad Request: Item already added / Failed to add to favorites"
// @Failure 401 {object} res.CommonRes "Unauthorized: Invalid or expired token"
// @Router /addToFavourite/{id} [post]
func (h *FavHandler) AddFavItem(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])
	dishId := c.Params("id")

	err := h.uc.AddFavItem(userId, dishId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully added to fav",
			})
	case e.ErrConflict:
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "item already added",
				Error:   err.Error(),
			})
	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to added fav",
				Error:   err.Error(),
			})
	}

}

// @Summary Delete item from favorites
// @Description Deletes a dish from the user's favorites
// @Security Bearer
// @Tags User Favorites
// @Accept json
// @Produce json
// @Param id path string true "Dish ID" 
// @Success 200 {object} res.CommonRes "Success: Dish deleted from favorites successfully"
// @Failure 400 {object} res.CommonRes "Bad Request: Item already deleted / Failed to delete from favorites"
// @Failure 401 {object} res.CommonRes "Unauthorized: Invalid or expired token"
// @Router /favourites/{id}/delete [delete]
func (h *FavHandler) DeleteFavItem(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])
	dishId := c.Params("id")

	err := h.uc.DeleteFavItem(userId, dishId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully deleted to fav",
			})
	case e.ErrNotFound:
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "item already deleted",
				Error:   err.Error(),
			})
	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to added fav",
				Error:   err.Error(),
			})
	}
}

// @Summary View favorite items
// @Description Fetches the list of user's favorite dishes
// @Security Bearer
// @Tags User Favorites
// @Accept json
// @Produce json
// @Success 200 {object} res.CommonRes "Success: Favorites fetched successfully"
// @Failure 200 {object} res.CommonRes "Success: Favorites are empty"
// @Failure 401 {object} res.CommonRes "Unauthorized: Invalid or expired token"
// @Router /favourites [get]

func (h *FavHandler) ViewFavItems(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	favList, err := h.uc.ViewFavourites(userId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully fetched favorites",
				Result:  favList,
			})
	case e.ErrIsEmpty:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "favourites are empty",
				Error:   err.Error(),
			})
	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to added fav",
				Error:   err.Error(),
			})
	}
}
