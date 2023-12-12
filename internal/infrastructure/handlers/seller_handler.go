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
			Token:   *token,
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
			Token:   *token,
		})
}

// @Summary		Get paginated list of sellers
// @Description	Retrieve a paginated list of sellers for the user
// @Tags			Common
// @Accept			json
// @Produce		json
// @Param			p	query		string				false	"Page number (default: 1)"
// @Param			l	query		string				false	"Number of items per page"
// @Success		200	{object}	res.SellerListRes	"Successfully fetched sellers"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/user/sellers [get]
func (h *SellerHandler) GetSellersPage(c *fiber.Ctx) error {
	page := c.Query("p", "1")
	limit := c.Query("l", "10")

	sellerList, err := h.usecase.GetSellersPage(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch sellers",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerListRes{
			Status:     "success",
			Message:    "successfully fetched seller",
			SellerList: *sellerList,
		})

}

// @Summary		Get a seller
// @Description	Retrieve a specific seller by ID for the user
// @Tags			Common
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Seller ID"	int
// @Success		200	{object}	res.SingleSellerRes	"Successfully fetched seller"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/user/sellers/{id} [get]
func (h *SellerHandler) GetSeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	seller, err := h.usecase.GetSeller(sellerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SingleSellerRes{
			Status:  "success",
			Message: "successfully fetched seller",
			Seller:  *seller,
		})
}

// @Summary		Get seller profile
// @Description	Retrieve a seller profile 
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.CommonRes	"Successfully fetched seller profile"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/seller/profile [get]
func (h *SellerHandler) GetSellerProfile(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully fetched seller profile",
			Result:  seller,
		})
}

// @Summary		Search sellers
// @Description	Search for sellers based on a query
// @Tags			Common
// @Accept			json
// @Produce		json
// @Param			q	query		string				true	"Search query"
// @Success		200	{object}	res.SellerListRes	"Successfully fetched sellers"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/search/sellers [get]
func (h *SellerHandler) SearchSeller(c *fiber.Ctx) error {
	searchQuery := c.Query("q")

	sellersList, err := h.usecase.SearchVerifiedSeller(searchQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch sellers",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerListRes{
			Status:     "success",
			Message:    "successfully fetched sellers",
			SellerList: *sellersList,
		})
}
