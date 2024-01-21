package usecases

import (
	"fmt"
	"strconv"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type favUseCase struct {
	favRepo   interfaces.IFavoriteRepository
	dishUcase i.IDishUseCase
}

func NewFavouriteUsecase(
	favRepo interfaces.IFavoriteRepository,
	dishUcase i.IDishUseCase,
) i.IFavouriteUseCase {
	return &favUseCase{favRepo, dishUcase}
}

func (uc *favUseCase) AddFavItem(userId, dishId string) error {
	uid, err := strconv.ParseUint(userId, 10, 0)
	if err != nil {
		return err
	}
	did, err := strconv.ParseUint(dishId, 10, 0)
	if err != nil {
		return err
	}

	var favItem = entities.Favourite{
		UserID: uint(uid),
		DishID: uint(did),
	}
	return uc.favRepo.Create(userId, dishId, &favItem)
}

func (uc *favUseCase) ViewFavourites(userId string) (*[]entities.Dish, error) {
	favList, err := uc.favRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	list := *favList
	var dishList []entities.Dish
	for i := range list {
		dish, err := uc.dishUcase.GetDish(fmt.Sprint(list[i].DishID))
		if err != nil && err != e.ErrNotFound{
			return nil, err
		}
		if err != e.ErrNotFound {
			dishList = append(dishList, *dish)
		}
	}
	return &dishList, nil
}

func (uc *favUseCase) DeleteFavItem(userId, dishId string) error {
	return uc.favRepo.Delete(userId, dishId)
}
