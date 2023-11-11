package handlers

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	imageuploader "github.com/abdullahnettoor/food-delivery-eCommerce/internal/services/image_uploader"
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

// @Summary		Create a dish
// @Description	Add a new dish for the seller
// @Security		Bearer
// @Tags			Seller
// @Accept			multipart/form-data
// @Produce		json
// @Param			image	formData	file			true				"Image file for the dish"
// @Param			req		body		formData		req.CreateDishReq	true	"Dish creation request"
// @Success		200		{object}	res.CommonRes	"Successfully created dish"
// @Failure		400		{object}	res.CommonRes	"Bad Request"
// @Failure		401		{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500		{object}	res.CommonRes	"Internal Server Error"
// @Router			/seller/addDish [post]
func (h *SellerHandler) CreateDish(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)
	sellerId := fmt.Sprint(seller["sellerId"])
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

	formFile, err := c.FormFile("image")
	if err != nil {
		fmt.Println("Error is", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to get image from form",
			})
	}

	wd, err := os.Getwd()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to get directory",
			})
	}
	path := filepath.Join(wd, "tmp", formFile.Filename)

	trgt, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to open file path in server",
			})
	}
	defer trgt.Close()

	f, err := formFile.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to open file from form",
			})
	}

	if _, err := io.Copy(trgt, f); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to copy form file to temp",
			})
	}

	ctx := context.Background()
	imgUploader := imageuploader.NewUploadImage()
	fileName := fmt.Sprintf("%s-%s", sellerId, strings.ToLower(strings.ReplaceAll(req.Name, " ", "-")))

	url, err := imgUploader.Handler(ctx, path, fileName, "dishes")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to upload file to cloud",
			})
	}

	req.ImageUrl = url

	e := os.Remove(path)
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to delete temp image",
			})
	}

	if err := h.usecase.AddDish(sellerId, &req); err != nil {
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

// @Summary		Update a dish
// @Description	Modify an existing dish by ID for the seller
// @Security		Bearer
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Dish ID"	int
// @Param			req	body		req.UpdateDishReq	true	"Dish update request"
// @Success		200	{object}	res.CommonRes		"Successfully updated dish"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/seller/dishes/{id} [put]
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

// @Summary		Get a dish
// @Description	Retrieve a specific dish by ID for the seller
// @Security		Bearer
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Dish ID"	int
// @Success		200	{object}	res.SingleDishRes	"Successfully fetched dish"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/seller/dishes/{id} [get]
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

// @Summary		Get all dishes
// @Description	Retrieve a list of all dishes for the seller
// @Security		Bearer
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.DishListRes	"Successfully fetched dishes"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/seller/dishes [get]
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

// @Summary		Delete a dish
// @Description	Remove a specific dish by ID for the seller
// @Security		Bearer
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"Dish ID"	int
// @Success		200	{object}	res.CommonRes	"Successfully deleted dish"
// @Failure		400	{object}	res.CommonRes	"Bad Request"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/seller/dishes/{id} [delete]
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
