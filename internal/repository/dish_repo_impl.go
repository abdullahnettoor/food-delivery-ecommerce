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

func (repo *DishRepository) FindPageWise(sellerId, categoryId string, page, limit uint) (*[]entities.Dish, error) {
	var dishList []entities.Dish

	query := ` SELECT * FROM dishes WHERE availability = true AND deleted = false`

	if sellerId != "" {
		query += " AND seller_id = " + sellerId
	}
	if categoryId != "" {
		query += " AND category_id = " + categoryId
	}

	query += " ORDER BY id DESC "

	offset := (page - 1) * uint(limit)
	if limit != 0 && page != 0 {
		query += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, limit)
	}

	res := repo.DB.Raw(query).Scan(&dishList)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
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

	res := repo.DB.Raw(query, id).Scan(&dish)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &dish, nil
}

func (repo *DishRepository) FindBySeller(sellerId, category_id string) (*[]entities.Dish, error) {
	var dishList []entities.Dish

	query := `
	SELECT * 
	FROM dishes
	WHERE deleted = false` + " AND seller_id = " + sellerId

	if category_id != "" {
		query += " AND category_id = " + category_id
	}

	query += " ORDER BY id DESC "

	res := repo.DB.Raw(query).Scan(&dishList)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
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

	res := repo.DB.Raw(query, id, sellerId).Scan(&dish)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
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

	query := fmt.Sprintf("UPDATE dishes SET name='%v', description = '%v', price = '%v', sale_price='%v', quantity = '%v', category_id = '%v', is_veg = '%v' , availability = '%v' WHERE id = '%v'",
		dish.Name, dish.Description, dish.Price, dish.SalePrice, dish.Quantity, dish.CategoryID, dish.IsVeg, dish.Availability, id)

	res := repo.DB.Exec(query)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
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

	res := repo.DB.Exec(query, id, sellerId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return e.ErrNotFound
	}

	return nil

}

func (repo *DishRepository) Search(search string) (*[]entities.Dish, error) {
	var dishList []entities.Dish

	query := fmt.Sprintf("SELECT * FROM dishes WHERE (name ILIKE '%%%s%%') OR (description ILIKE '%%%s%%') AND deleted = false 	ORDER BY id DESC", search, search)

	fmt.Println("Query is", query)

	res := repo.DB.Raw(query).Scan(&dishList)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &dishList, nil
}

func (repo *DishRepository) ReduceStock(id string, quantity uint) error {
	return repo.DB.Exec(`
	UPDATE dishes
	SET quantity = quantity - ?
	WHERE id = ?`,
		quantity, id).Error
}

func (repo *DishRepository) IncreaseStock(id string, quantity uint) error {
	return repo.DB.Exec(`
	UPDATE dishes
	SET quantity = quantity + ?
	WHERE id = ?`,
		quantity, id).Error
}
