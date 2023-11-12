package repository

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.ICategoryRepository {
	return &CategoryRepository{db}
}

func (repo *CategoryRepository) FindByID(id string) (*entities.Category, error) {
	var category entities.Category

	query := `
	SELECT * 
	FROM categories 
	WHERE id = ?`

	if err := repo.DB.Raw(query, id).Scan(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *CategoryRepository) FindAll() (*[]entities.Category, error) {
	var categoryList []entities.Category

	query := `SELECT * FROM categories`

	if err := repo.DB.Raw(query).Scan(&categoryList).Error; err != nil {
		return nil, err
	}
	return &categoryList, nil
}

func (repo *CategoryRepository) Create(category *entities.Category) error {

	query := repo.DB.Raw(`
	SELECT * 
	FROM categories
	WHERE name ILIKE ?`, category.Name).Scan(category)

	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected != 0 {
		return e.ErrConflict
	}

	return repo.DB.Create(&category).Error
}

func (repo *CategoryRepository) Update(id, name string) error {
	var category entities.Category

	query := repo.DB.Raw(`
	SELECT * 
	FROM categories
	WHERE name ILIKE ?`, name).Scan(&category)

	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected != 0 {
		return e.ErrConflict
	}

	err := repo.DB.Exec(`
	UPDATE categories
	SET name = ?
	WHERE id = ?`, name, id).Error

	return err
}
