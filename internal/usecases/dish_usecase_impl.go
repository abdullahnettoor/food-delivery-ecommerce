package usecases

import (
	"strconv"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type dishUsecase struct {
	dishRepo   interfaces.IDishRepository
	offerUcase i.IOfferUseCase
}

func NewDishUsecase(dishRepo interfaces.IDishRepository, offerUcase i.IOfferUseCase) i.IDishUseCase {
	return &dishUsecase{dishRepo, offerUcase}
}

func (uc *dishUsecase) AddDish(sellerId string, req *req.CreateDishReq) error {

	id, err := strconv.ParseUint(sellerId, 10, 0)
	if err != nil {
		return err
	}

	newDish := entities.Dish{
		SellerID:     uint(id),
		Name:         req.Name,
		Description:  req.Description,
		ImageUrl:     req.ImageUrl,
		Price:        req.Price,
		SalePrice:    req.SalePrice,
		Quantity:     req.Quantity,
		CategoryID:   req.CategoryID,
		IsVeg:        req.IsVeg,
		Availability: req.Availability,
	}
	if req.SalePrice == 0 {
		newDish.SalePrice = req.Price
	}
	return uc.dishRepo.Create(&newDish)
}

func (uc *dishUsecase) UpdateDish(dishId, sellerId string, req *req.UpdateDishReq) (*entities.Dish, error) {

	id, err := strconv.ParseUint(sellerId, 10, 0)
	if err != nil {
		return nil, err
	}

	updatedDish := entities.Dish{
		SellerID:     uint(id),
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		SalePrice:    req.SalePrice,
		Quantity:     req.Quantity,
		CategoryID:   req.CategoryID,
		IsVeg:        req.IsVeg,
		Availability: req.Availability,
	}
	if req.SalePrice == 0 {
		updatedDish.SalePrice = req.Price
	}
	return uc.dishRepo.Update(dishId, &updatedDish)
}

func (uc *dishUsecase) GetAllDishesBySeller(sellerId, category_id string) (*[]entities.Dish, error) {
	dishList, err := uc.dishRepo.FindBySeller(sellerId, category_id)
	if err != nil {
		return nil, err
	}
	return uc.offerUcase.ApplyOfferToDishList(dishList)
}

func (uc *dishUsecase) GetDishBySeller(id, sellerId string) (*entities.Dish, error) {
	dish, err := uc.dishRepo.FindBySellerAndID(id, sellerId)
	if err != nil {
		return nil, err
	}
	return uc.offerUcase.ApplyOfferToDish(dish)
}

func (uc *dishUsecase) DeleteDish(id, sellerId string) error {
	return uc.dishRepo.Delete(id, sellerId)
}

func (uc *dishUsecase) SearchDish(search string) (*[]entities.Dish, error) {
	dishList, err := uc.dishRepo.Search(search)
	if err != nil {
		return nil, err
	}
	return uc.offerUcase.ApplyOfferToDishList(dishList)
}

func (uc *dishUsecase) GetDishesPage(sellerId, categoryId string, page, limit string) (*[]entities.Dish, error) {
	p, err := strconv.ParseUint(page, 10, 0)
	if err != nil {
		return nil, err
	}
	l, err := strconv.ParseUint(limit, 10, 0)
	if err != nil {
		return nil, err
	}

	dishList, err := uc.dishRepo.FindPageWise(sellerId, categoryId, uint(p), uint(l))
	if err != nil {
		return nil, err
	}

	return uc.offerUcase.ApplyOfferToDishList(dishList)
}

func (uc *dishUsecase) GetDish(id string) (*entities.Dish, error) {
	dish, err := uc.dishRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return uc.offerUcase.ApplyOfferToDish(dish)
}

func (uc *dishUsecase) ReduceStock(id string, quantity uint) error {
	return uc.dishRepo.ReduceStock(id, quantity)
}
