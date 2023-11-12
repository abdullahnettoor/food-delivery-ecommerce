package handlers

import (
	"fmt"

	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	usecase interfaces.ICartUseCase
}

func NewCartHandler(uc interfaces.ICartUseCase) *CartHandler {
	return &CartHandler{uc}
}

//	@Summary		View Cart
//	@Description	Fetch the user's cart
//	@Security		Bearer
//	@Tags			Cart
//	@Produce		json
//	@Success		200	{object}	res.ViewCartRes	"Successfully fetched cart"
//	@Success		200	{object}	res.CommonRes	"Cart is empty"
//	@Failure		400	{object}	res.CommonRes	"Bad Request"
//	@Failure		401	{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes	"Internal Server Error"
//	@Router			/cart [get]
func (h *CartHandler) ViewCart(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)

	cartId := fmt.Sprint(user["userId"])

	cart, err := h.usecase.ViewCart(cartId)
	if err != nil && err != e.ErrNotFound {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch cart",
				Error:   err.Error(),
			})
	}
	if err == e.ErrNotFound {
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Message: "cart is empty",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.ViewCartRes{
			Status:  "success",
			Message: "successfully fetched cart",
			Cart:    *cart,
		})
}

//	@Summary		Add to Cart
//	@Description	Add a dish to the user's cart
//	@Security		Bearer
//	@Tags			Cart
//	@Produce		json
//	@Param			id	path		string			true	"Dish ID"
//	@Success		200	{object}	res.CommonRes	"Successfully added to cart"
//	@Failure		400	{object}	res.CommonRes	"Bad Request"
//	@Failure		401	{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes	"Internal Server Error"
//	@Router			/addToCart/{id} [post]
func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)

	cartId := fmt.Sprint(user["userId"])
	dishId := c.Params("id")

	err := h.usecase.AddtoCart(cartId, dishId)
	if err != nil && err != e.ErrNotAvailable {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to add item to cart",
			})
	}
	if err == e.ErrNotAvailable {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "selected item is not available",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "suucessfully added to cart",
		})

}

//	@Summary		Decrement Cart Item
//	@Description	Decrement the quantity of a dish in the user's cart
//	@Security		Bearer
//	@Tags			Cart
//	@Produce		json
//	@Param			id	path		string			true	"Dish ID"
//	@Success		200	{object}	res.CommonRes	"Successfully decremented cart item"
//	@Failure		400	{object}	res.CommonRes	"Bad Request"
//	@Failure		401	{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes	"Internal Server Error"
//	@Router			/cart/{id}/decrement [patch]
func (h *CartHandler) DecrementCartItem(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)

	cartId := fmt.Sprint(user["userId"])
	dishId := c.Params("id")

	err := h.usecase.DecrementCartItem(cartId, dishId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to decrement item in cart",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "suucessfully decremented to cart item",
		})
}

//	@Summary		Delete Cart Item
//	@Description	Delete a dish from the user's cart
//	@Security		Bearer
//	@Tags			Cart
//	@Produce		json
//	@Param			id	path		string			true	"Dish ID"
//	@Success		200	{object}	res.CommonRes	"Successfully deleted cart item"
//	@Failure		400	{object}	res.CommonRes	"Bad Request"
//	@Failure		401	{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes	"Internal Server Error"
//	@Router			/cart/{id}/deleteItem [delete]
func (h *CartHandler) DeleteCartItem(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)

	cartId := fmt.Sprint(user["userId"])
	dishId := c.Params("id")

	err := h.usecase.DeleteCartItem(cartId, dishId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to delete cart item",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "suucessfully deleted cart item",
		})
}

//	@Summary		Empty Cart
//	@Description	Empty the user's cart
//	@Security		Bearer
//	@Tags			Cart
//	@Produce		json
//	@Success		200	{object}	res.CommonRes	"Successfully emptied cart"
//	@Failure		400	{object}	res.CommonRes	"Bad Request"
//	@Failure		401	{object}	res.CommonRes	"Unauthorized Access"
//	@Failure		500	{object}	res.CommonRes	"Internal Server Error"
//	@Router			/cart/empty [delete]
func (h *CartHandler) EmptyCart(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)

	cartId := fmt.Sprint(user["userId"])

	err := h.usecase.EmptyCart(cartId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to empty cart",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "suucessfully emptied cart",
		})
}
