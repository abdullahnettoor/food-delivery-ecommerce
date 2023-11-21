package usecases

import (
	"errors"
	"fmt"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type cartUseCase struct {
	cartRepo interfaces.ICartRepository
	dishUcase i.IDishUseCase
}

func NewCartUsecase(
	cartRepo interfaces.ICartRepository,
	dishUcase i.IDishUseCase,
) i.ICartUseCase {
	return &cartUseCase{cartRepo, dishUcase}
}

func (uc *cartUseCase) AddtoCart(id, dishId string) error {

	_, cartErr := uc.cartRepo.FindCart(id)
	if cartErr != nil && cartErr != e.ErrNotFound {
		return cartErr
	}

	dish, err := uc.dishUcase.GetDish(dishId)
	if err != nil {
		return err
	}

	sellerId := fmt.Sprint(dish.SellerID)
	if cartErr == e.ErrNotFound {
		if err := uc.cartRepo.CreateCart(id, sellerId); err != nil {
			return err
		}
	}

	cart, cartErr := uc.cartRepo.FindCart(id)
	if cartErr != nil && cartErr != e.ErrNotFound {
		return cartErr
	}
	if cart.SellerID != dish.SellerID {
		return errors.New("can't add items from different seller")
	}
	if !dish.Availability || dish.Quantity == 0 {
		return e.ErrNotAvailable
	}

	return uc.cartRepo.AddToCart(id, dishId)
}

func (uc *cartUseCase) ViewCart(id string) (*entities.Cart, error) {

	cart, err := uc.cartRepo.FindCart(id)
	if err != nil {
		return nil, err
	}

	items, err := uc.cartRepo.FindCartItems(id)
	if err != nil {
		return nil, err
	}

	var cartItems []entities.CartItem
	for _, item := range *items {
		dish, err := uc.dishUcase.GetDish(fmt.Sprint(item.DishID))
		if err != nil {
			return nil, err
		}

		item = entities.CartItem{
			Quantity: item.Quantity,
			Dish: entities.Dish{
				ID:           dish.ID,
				SellerID:     dish.SellerID,
				Name:         dish.Name,
				Description:  dish.Description,
				Price:        dish.Price,
				SalePrice:    dish.SalePrice,
				ImageUrl:     dish.ImageUrl,
				CategoryID:   dish.CategoryID,
				IsVeg:        dish.IsVeg,
				Availability: dish.Availability,
			},
		}
		cartItems = append(cartItems, item)
		fmt.Println("\nCart Item is", item)
	}

	cart.CartItems = cartItems

	return cart, nil
}

func (uc *cartUseCase) DeleteCartItem(id, dishId string) error {
	return uc.cartRepo.DeleteItem(id, dishId)
}

func (uc *cartUseCase) DecrementCartItem(id, dishId string) error {
	return uc.cartRepo.DecrementItem(id, dishId)
}

func (uc *cartUseCase) EmptyCart(id string) error {
	return uc.cartRepo.DeleteCart(id)
}
