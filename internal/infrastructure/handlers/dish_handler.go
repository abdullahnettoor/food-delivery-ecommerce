package handlers

import (
	"context"
	"fmt"
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

type DishHandler struct {
	dishUc interfaces.IDishUseCase
}

func NewDishHandler(uc interfaces.IDishUseCase) *DishHandler {
	return &DishHandler{uc}
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
func (h *DishHandler) CreateDish(c *fiber.Ctx) error {
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

	path := filepath.Join(os.TempDir(), formFile.Filename)

	if err := c.SaveFile(formFile, path); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to open file path in server",
			})
	}

	file, err := os.Open(path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to open file path in server",
			})
	}

	ctx := context.Background()
	imgUploader := imageuploader.NewUploadImage()
	fileName := fmt.Sprintf(
		"%s-%s",
		sellerId,
		strings.ToLower(
			strings.ReplaceAll(req.Name, " ", "-")))

	url, err := imgUploader.Handler(ctx, fileName, "dishes", file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to upload file to cloud",
			})
	}
	file.Close()

	req.ImageUrl = url

	err = os.Remove(path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to delete temp image",
			})
	}

	if err := h.dishUc.AddDish(sellerId, &req); err != nil {
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
func (h *DishHandler) UpdateDish(c *fiber.Ctx) error {
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

	dish, err := h.dishUc.UpdateDish(dishId, fmt.Sprint(seller["sellerId"]), &req)
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
func (h *DishHandler) GetDishBySeller(c *fiber.Ctx) error {
	dishId := c.Params("id")
	seller := c.Locals("SellerModel").(map[string]any)

	dish, err := h.dishUc.GetDishBySeller(dishId, fmt.Sprint(seller["sellerId"]))
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
// @Param			category	query		int				false	"Category Id"
// @Success		200			{object}	res.DishListRes	"Successfully fetched dishes"
// @Failure		401			{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500			{object}	res.CommonRes	"Internal Server Error"
// @Router			/seller/dishes [get]
func (h *DishHandler) GetAllDishBySeller(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)
	sellerId := fmt.Sprint(seller["sellerId"])
	categoryId := c.Query("category")

	dishList, err := h.dishUc.GetAllDishesBySeller(sellerId, categoryId)
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
func (h *DishHandler) DeleteDish(c *fiber.Ctx) error {
	dishId := c.Params("id")
	seller := c.Locals("SellerModel").(map[string]any)

	err := h.dishUc.DeleteDish(dishId, fmt.Sprint(seller["sellerId"]))
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

// @Summary		Get paginated list of dishes
// @Description	Retrieve a paginated list of dishes for the user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			p	query		string			false	"Page number (default: 1)"
// @Param			l	query		string			false	"Number of items per page"
// @Param			category	query		string			false	"Item category"
// @Success		200	{object}	res.DishListRes	"Successfully fetched dishes"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/dishes [get]
func (h *DishHandler) GetDishesPage(c *fiber.Ctx) error {
	page := c.Query("p", "1")
	limit := c.Query("l")
	categoryId := c.Query("category")

	dishList, err := h.dishUc.GetDishesPage(categoryId, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch dishes",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.DishListRes{
			Status:   "success",
			Message:  "successfully fetched dish",
			DishList: *dishList,
		})

}

// @Summary		Get a dish
// @Description	Retrieve a specific dish by ID for the user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Dish ID"	int
// @Success		200	{object}	res.SingleDishRes	"Successfully fetched dish"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/dishes/{id} [get]
func (h *DishHandler) GetDish(c *fiber.Ctx) error {
	dishId := c.Params("id")

	dish, err := h.dishUc.GetDish(dishId)
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

// @Summary		Search dishes
// @Description	Search for dishes based on a query
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			q	query		string			true	"Search query"
// @Success		200	{object}	res.DishListRes	"Successfully fetched dishes"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/search/dishes [get]
func (h *DishHandler) SearchDish(c *fiber.Ctx) error {
	searchQuery := c.Query("q")

	dishList, err := h.dishUc.SearchDish(searchQuery)
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
			Message:  "successfully fetched dish",
			DishList: *dishList,
		})
}
