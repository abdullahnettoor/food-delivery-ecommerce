package usecases

import (
	"strconv"
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type offerUsecase struct {
	repo interfaces.IOfferRepository
}

func NewOfferUsecase(repo interfaces.IOfferRepository) i.IOfferUseCase {
	return &offerUsecase{repo}
}

func (uc *offerUsecase) CreateOffer(sellerId string, req *req.CreateOfferReq) error {
	sId, err := strconv.ParseUint(sellerId, 10, 0)
	if err != nil {
		return err
	}

	offer := entities.CategoryOffer{
		Title:      req.Title,
		SellerID:   uint(sId),
		CategoryID: req.CategoryID,
		Percentage: req.Percentage,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
	}
	if req.StartDate.IsZero() {
		offer.StartDate = time.Now()
	}
	if req.EndDate.IsZero() {
		offer.EndDate = time.Now().Add(time.Hour * 730)
	}

	return uc.repo.Create(&offer)
}

func (uc *offerUsecase) UpdateOffer(id, sellerId string, req *req.UpdateOfferReq) error {
	oId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return err
	}
	sId, err := strconv.ParseUint(sellerId, 10, 0)
	if err != nil {
		return err
	}

	var updatedOffer = entities.CategoryOffer{
		ID:         uint(oId),
		Title:      req.Title,
		SellerID:   uint(sId),
		CategoryID: req.CategoryID,
		Percentage: req.Percentage,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
	}
	if req.StartDate.IsZero() {
		updatedOffer.StartDate = time.Now()
	}
	if req.EndDate.IsZero() {
		updatedOffer.EndDate = time.Now().Add(time.Hour * 730)
	}
	if req.Status == "Closed" {
		updatedOffer.EndDate = time.Now()
	}

	return uc.repo.Update(id, &updatedOffer)
}

func (uc *offerUsecase) UpdateOfferStatus(id, status string) error {
	return uc.repo.UpdateStatus(id, status)
}

func (uc *offerUsecase) GetAllOffer() (*[]entities.CategoryOffer, error) {
	return uc.repo.FindAll()
}

func (uc *offerUsecase) GetOffersBySeller(sellerId string) (*[]entities.CategoryOffer, error) {
	return uc.repo.FindAllForSeller(sellerId)
}
