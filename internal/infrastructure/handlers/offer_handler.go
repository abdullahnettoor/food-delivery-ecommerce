package handlers

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	imageuploader "github.com/abdullahnettoor/food-delivery-eCommerce/internal/services/image_uploader"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
)

type OfferHandler struct {
	uc interfaces.IOfferUseCase
}

func NewOfferHandler(uc interfaces.IOfferUseCase) *OfferHandler {
	return &OfferHandler{uc}
}

//	@Summary		Get all offers
//	@Description	Fetches a list of all offers
//	@Tags			Common
//	@Produce		json
//	@Success		200	{object}	res.OfferListRes	"Success: List of offers fetched successfully"
//	@Failure		500	{object}	res.CommonRes		"Internal Server Error: Failed to fetch offers"
//	@Router			/offers [get]
func (h *OfferHandler) GetAllOffers(c *fiber.Ctx) error {
	offerList, err := h.uc.GetAllOffer()
	if err != nil && err != e.ErrIsEmpty {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch offers",
			})
	}

	if err == e.ErrIsEmpty {
		return c.Status(fiber.StatusOK).
			JSON(res.CommonRes{
				Status:  "success",
				Error:   err.Error(),
				Message: "there are no offers now",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.OfferListRes{
			Status:    "success",
			Message:   "offers fetched successfully",
			OfferList: *offerList,
		})
}

//	@Summary		Get offers by seller
//	@Description	Fetches a list of offers associated with the seller
//	@Security		Bearer
//	@Tags			Seller Offer
//	@Produce		json
//	@Success		200	{object}	res.OfferListRes	"Success: List of offers fetched successfully"
//	@Failure		500	{object}	res.CommonRes		"Internal Server Error: Failed to fetch offers"
//	@Router			/seller/offers [get]
func (h *OfferHandler) GetOffersBySeller(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)
	sellerId := fmt.Sprint(seller["sellerId"])

	offerList, err := h.uc.GetOffersBySeller(sellerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch offers",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.OfferListRes{
			Status:    "success",
			Message:   "offers fetched successfully",
			OfferList: *offerList,
		})
}

//	@Summary		Create an offer
//	@Description	Create new offer for the seller
//	@Security		Bearer
//	@Tags			Seller Offer
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			image	formData	file			true				"Image file for the dish"
//	@Param			req		body		formData		req.CreateOfferReq	true	"Create Offer Request"
//	@Success		200		{object}	res.CommonRes	"Success: Offer updated successfully"
//	@Failure		400		{object}	res.CommonRes	"Bad Request: Invalid inputs"
//	@Failure		500		{object}	res.CommonRes	"Internal Server Error: Error occurred while updating offer"
//	@Router			/seller/offers/addOffer [post]
func (h *OfferHandler) CreateOffer(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)
	sellerId := fmt.Sprint(seller["sellerId"])

	var req req.CreateOfferReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if errs := requestvalidation.ValidateRequest(req); errs != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(errs),
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
	fmt.Println("File Path is", path)

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
		"%v-%v",
		time.Now().Unix(),
		sellerId)

	fmt.Println("File is", *file)

	url, err := imgUploader.Handler(ctx, fileName, "offers", file)
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

	if err := h.uc.CreateOffer(sellerId, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "error occured while creating offer",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "offer created successfully",
		})
}

//	@Summary		Update an offer
//	@Description	Updates details of a specific offer for the seller
//	@Security		Bearer
//	@Tags			Seller Offer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Offer ID"
//	@Param			req	body		req.UpdateOfferReq	true	"Update Offer Request"
//	@Success		200	{object}	res.CommonRes		"Success: Offer updated successfully"
//	@Failure		400	{object}	res.CommonRes		"Bad Request: Invalid inputs"
//	@Failure		500	{object}	res.CommonRes		"Internal Server Error: Error occurred while updating offer"
//	@Router			/seller/offers/{id} [put]
func (h *OfferHandler) UpdateOffer(c *fiber.Ctx) error {
	seller := c.Locals("SellerModel").(map[string]any)
	id := c.Params("id")
	sellerId := fmt.Sprint(seller["sellerId"])

	var req req.UpdateOfferReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if errs := requestvalidation.ValidateRequest(req); errs != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(errs),
				Message: "invalid inputs",
			})
	}

	if err := h.uc.UpdateOffer(id, sellerId, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "error occured while updating offer",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "offer updated successfully",
		})
}

//	@Summary		Update offer status
//	@Description	Updates the status of a specific offer
//	@Security		Bearer
//	@Tags			Seller Order
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Offer ID"
//	@Param			status	query		string			true	"New offer status"
//	@Success		200		{object}	res.CommonRes	"Success: Offer status updated successfully"
//	@Failure		500		{object}	res.CommonRes	"Internal Server Error: Error occurred while updating offer status"
//	@Router			/seller/offers/{id} [patch]
func (h *OfferHandler) UpdateOfferStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Query("status")

	if err := h.uc.UpdateOfferStatus(id, status); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "error occured while updating offer status",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "offer status updated successfully",
		})
}
