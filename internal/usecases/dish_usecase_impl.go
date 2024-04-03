package usecases

import (
	"fmt"
	"strconv"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type dishUsecase struct {
	dishRepo  interfaces.IDishRepository
	offerRepo interfaces.IOfferRepository
}

func NewDishUsecase(dishRepo interfaces.IDishRepository, offerRepo interfaces.IOfferRepository) i.IDishUseCase {
	return &dishUsecase{dishRepo, offerRepo}
}

func (uc *dishUsecase) ApplyOfferToDishList(dishList *[]entities.Dish) (*[]entities.Dish, error) {
	list := *dishList

	for i := range list {
		cId := fmt.Sprint(list[i].CategoryID)
		sId := fmt.Sprint(list[i].SellerID)

		offer, err := uc.offerRepo.FindBySellerAndCategory(sId, cId)
		if err != nil && err != e.ErrNotFound {
			fmt.Println("Error is", err)
			return nil, err
		}
		list[i].SalePrice = list[i].Price - (list[i].Price * float64(offer.Percentage) / 100)
	}

	return &list, nil
}

func (uc *dishUsecase) ApplyOfferToDish(dish *entities.Dish) (*entities.Dish, error) {

	sId := fmt.Sprint(dish.SellerID)
	cId := fmt.Sprint(dish.CategoryID)
	offer, err := uc.offerRepo.FindBySellerAndCategory(sId, cId)
	if err != nil && err != e.ErrNotFound {
		fmt.Println("Error is", err)
		return nil, err
	}

	dish.SalePrice = dish.Price - (dish.Price * float64(offer.Percentage) / 100)
	fmt.Printf("Offer Applied for %v \nRegular Price: %v\nSale Price: %v", dish.Name, dish.Price, dish.SalePrice)

	return dish, nil
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
	return uc.ApplyOfferToDishList(dishList)
}

func (uc *dishUsecase) GetDishBySeller(id, sellerId string) (*entities.Dish, error) {
	dish, err := uc.dishRepo.FindBySellerAndID(id, sellerId)
	if err != nil {
		return nil, err
	}
	return uc.ApplyOfferToDish(dish)
}

func (uc *dishUsecase) DeleteDish(id, sellerId string) error {
	return uc.dishRepo.Delete(id, sellerId)
}

func (uc *dishUsecase) SearchDish(search, sellerId string) (*[]entities.Dish, error) {
	dishList, err := uc.dishRepo.Search(search, sellerId)
	if err != nil {
		return nil, err
	}
	return uc.ApplyOfferToDishList(dishList)
}

func (uc *dishUsecase) GetDishesPage(sellerId, categoryId string, page, limit string) (*[]entities.Dish, error) {
	p, err := strconv.ParseUint(page, 10, 0)
	if err != nil {
		fmt.Println("Parsing Error 1", err)
		return nil, err
	}
	l, err := strconv.ParseUint(limit, 10, 0)
	if err != nil {
		fmt.Println("Parsing Error 2", err)
		return nil, err
	}

	dishList, err := uc.dishRepo.FindPageWise(sellerId, categoryId, uint(p), uint(l))
	if err != nil {
		fmt.Println("Error Finding Page wise", err)
		return nil, err
	}
	return uc.ApplyOfferToDishList(dishList)
}

func (uc *dishUsecase) GetDish(id string) (*entities.Dish, error) {
	dish, err := uc.dishRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return uc.ApplyOfferToDish(dish)
}

func (uc *dishUsecase) ReduceStock(id string, quantity uint) error {
	return uc.dishRepo.ReduceStock(id, quantity)
}
