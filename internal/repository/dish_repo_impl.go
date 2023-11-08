package repository

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type DishRepository struct {
	DB *gorm.DB
}

func NewDishRepository(db *gorm.DB) interfaces.IDishRepository {
	return &DishRepository{db}
}

func (repo *DishRepository) FindPageWise(page int) (*[]entities.Dish, error) {
	var dishList []entities.Dish

	query := `
	SELECT *
	FROM dishes
	WHERE is_available = true
	AND deleted = false`

	if err := repo.DB.Raw(query).Scan(&dishList).Error; err != nil {
		return nil, err
	}

	return &dishList, nil
}

func (repo *DishRepository) FindByID(id string) (*entities.Dish, error) {
	var dish entities.Dish

	query := `
	SELECT *
	FROM dishes
	WHERE deleted = false
	AND id = ?`

	if err := repo.DB.Raw(query, id).Scan(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

func (repo *DishRepository) FindBySeller(sellerId string) (*[]entities.Dish, error) {
	var dishList []entities.Dish

	query := `
	SELECT * 
	FROM dishes
	WHERE seller_id = ?
	AND deleted = false`

	err := repo.DB.Raw(query, sellerId).Scan(&dishList).Error
	if err != nil {
		return nil, err
	}

	return &dishList, nil
}

func (repo *DishRepository) FindBySellerAndID(id, sellerId string) (*entities.Dish, error) {
	var dish entities.Dish

	query := `
	SELECT * 
	FROM dishes
	WHERE id = ?
	AND seller_id = ?
	AND deleted = false`

	err := repo.DB.Raw(query, id, sellerId).Scan(&dish).Error
	if err != nil {
		return nil, err
	}

	return &dish, nil
}

func (repo *DishRepository) Create(dish *entities.Dish) error {

	var dbDish entities.Dish

	query := repo.DB.Raw(`
	SELECT name 
	FROM dishes
	WHERE name = ?
	AND seller_id = ?
	`, dish.Name, dish.SellerID).Scan(&dbDish)

	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected != 0 {
		return e.ErrConflict
	}

	return repo.DB.Create(dish).Error
}

func (repo *DishRepository) Update(id string, dish *entities.Dish) (*entities.Dish, error) {
	var updatedDish entities.Dish

	q := repo.DB.Raw(`
	SELECT name 
	FROM dishes
	WHERE name = ?
	AND seller_id = ?
	AND id <> ?
	`, dish.Name, dish.SellerID, id).Scan(&updatedDish)

	if q.Error != nil {
		return nil, q.Error
	}
	if q.RowsAffected != 0 {
		return nil, e.ErrConflict
	}

	query := fmt.Sprintf("UPDATE dishes SET name='%v', description = '%v', price = '%v', quantity = '%v', category_id = '%v', is_veg = '%v' , availability = '%v' WHERE id = '%v'",
		dish.Name, dish.Description, dish.Price, dish.Quantity, dish.CategoryID, dish.IsVeg, dish.Availability, id)

	err := repo.DB.Exec(query).Error
	if err != nil {
		return nil, err
	}

	if err := repo.DB.Raw(`
	SELECT * 
	FROM dishes 
	WHERE id = ?`, id).
		Scan(&updatedDish).Error; err != nil {
		return nil, err
	}
	return &updatedDish, nil
}

func (repo *DishRepository) Delete(id, sellerId string) error {

	query := `
	UPDATE dishes
	SET deleted = true
	WHERE id = ?
	AND seller_id = ?`

	return repo.DB.Exec(query, id, sellerId).Error

}
