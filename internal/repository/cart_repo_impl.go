package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.ICartRepository {
	return &cartRepository{db}
}

func (repo *cartRepository) FindCart(id string) (*entities.Cart, error) {
	var cart entities.Cart

	res := repo.DB.Raw(`
	SELECT *
	FROM carts
	WHERE id = ?`, id).Scan(&cart)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &cart, nil
}

func (repo *cartRepository) FindCartItems(id string) (*[]entities.CartItem, error) {
	var cartItems []entities.CartItem

	res := repo.DB.Raw(`
	SELECT *
	FROM cart_items
	WHERE cart_id = ?`,
		id).Scan(&cartItems)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &cartItems, nil
}

func (repo *cartRepository) AddToCart(id, dishId string) error {
	var cart entities.Cart
	var cartItems []entities.CartItem

	res := repo.DB.Raw(`
	SELECT *
	FROM carts
	WHERE id = ?`,
		id).Scan(&cart)

	if res.Error != nil {
		return res.Error
	}

	if err := repo.DB.Raw(`
	SELECT *
	FROM cart_items
	WHERE cart_id = ?`,
		id).Scan(&cartItems).Error; err != nil {
		return err
	}

	for _, item := range cartItems {
		fmt.Println("Cart Item is", item)

		if fmt.Sprint(item.DishID) == dishId {
			item.Quantity += 1

			err := repo.DB.Exec(`
			UPDATE cart_items
			SET quantity = quantity + 1
			WHERE cart_id = ?
			AND dish_id = ?`,
				cart.ID, dishId).Error

			if err != nil {
				return err
			}
			return nil
		}
	}
	fmt.Println("Hello-----------------------")

	dId, _ := strconv.ParseUint(dishId, 10, 0)
	cId, _ := strconv.ParseUint(id, 10, 0)
	cartItem := entities.CartItem{
		CartID:   uint(cId),
		DishID:   uint(dId),
		Quantity: 1,
	}

	if err := repo.DB.Create(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

func (repo *cartRepository) DeleteItem(id, dishId string) error {

	if err := repo.DB.Exec(`
	DELETE FROM cart_items 
	WHERE cart_id = ? 
	AND dish_id = ?`,
		id, dishId).Error; err != nil {
		return err
	}

	if err := repo.DB.Exec(`
	DELETE FROM carts
	WHERE NOT EXISTS (
		SELECT 1 FROM cart_items
		WHERE cart_items.cart_id = ?
		)
	AND carts.id = ?`,
		id, id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *cartRepository) DecrementItem(id, dishId string) error {
	var cartItem entities.CartItem

	res := repo.DB.Raw(`
	SELECT * 
	FROM cart_items 
	WHERE cart_id = ? 
	AND dish_id = ?`,
		id, dishId).Scan(&cartItem)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("already deleted")
	}

	if cartItem.Quantity == 1 {
		return repo.DeleteItem(id, dishId)
	}

	cartItem.Quantity -= 1
	if err := repo.DB.Save(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

func (repo *cartRepository) DeleteCart(id string) error {
	return repo.DB.Exec(`
	DELETE FROM carts
	WHERE id = ?`,
		id).Error
}

func (repo *cartRepository) CreateCart(id, sellerId string) error {
	cartId, _ := strconv.ParseUint(id, 10, 0)
	sId, _ := strconv.ParseUint(sellerId, 10, 0)

	cart := entities.Cart{
		ID:       uint(cartId),
		SellerID: uint(sId),
	}

	return repo.DB.Save(&cart).Error
}
