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

// @Summary		Add item to favourites
// @Description	Adds a dish to the user's favourites
// @Security		Bearer
// @Tags			User Favourites
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Dish ID"
// @Success		200	{object}	res.CommonRes	"Success: Dish added to favourites successfully"
// @Failure		400	{object}	res.CommonRes	"Bad Request: Item already added / Failed to add to favourites"
// @Failure		401	{object}	res.CommonRes	"Unauthorized: Invalid or expired token"
// @Router			/addToFavourite/{id} [post]
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

// @Summary		Delete item from favourites
// @Description	Deletes a dish from the user's favourites
// @Security		Bearer
// @Tags			User Favourites
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Dish ID"
// @Success		200	{object}	res.CommonRes	"Success: Dish deleted from favourites successfully"
// @Failure		400	{object}	res.CommonRes	"Bad Request: Item already deleted / Failed to delete from favourites"
// @Failure		401	{object}	res.CommonRes	"Unauthorized: Invalid or expired token"
// @Router			/favourites/{id}/delete [delete]
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

// @Summary		View favourite items
// @Description	Fetches the list of user's favourite dishes
// @Security		Bearer
// @Tags			User Favourites
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.CommonRes	"Success: Favourites fetched successfully"
// @Failure		200	{object}	res.CommonRes	"Success: Favourites are empty"
// @Failure		401	{object}	res.CommonRes	"Unauthorized: Invalid or expired token"
// @Router			/favourites [get]
func (h *FavHandler) ViewFavItems(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	favList, err := h.uc.ViewFavourites(userId)

	switch err {
	case nil:
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "successfully fetched favourites",
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
