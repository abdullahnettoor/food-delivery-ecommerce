package repository

import (
	"errors"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.IUserRepository {
	return &UserRepository{DB: DB}
}

func (repo *UserRepository) FindAll() (*[]entities.User, error) {
	var usersList []entities.User

	if err := repo.DB.Raw(`
	SELECT *
	FROM users`).
		Scan(&usersList).Error; err != nil {
		return nil, err
	}

	return &usersList, nil
}

func (repo *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var userDetails entities.User
	query := repo.DB.Raw(`
		SELECT * 
		FROM users 
		WHERE email = ?`,
		email).Scan(&userDetails)

	if query.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return &userDetails, nil
}

func (repo *UserRepository) FindByID(id string) (*entities.User, error) {
	var userDetails entities.User
	query := repo.DB.Raw(`
		SELECT * 
		FROM users 
		WHERE id = ?`,
		id).Scan(&userDetails)

	if query.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return &userDetails, nil
}

func (repo *UserRepository) Create(user *entities.User) (*entities.User, error) {
	var newUser entities.User
	query := repo.DB.Raw(`
		SELECT * 
		FROM users 
		WHERE email = ?`,
		user.Email)

	if query.Error != nil {
		return nil, query.Error
	}

	if err := repo.DB.Create(&user).Scan(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (repo *UserRepository) Update(id string, user *entities.User) (*entities.User, error) {
	var newUser entities.User
	query := repo.DB.Raw(`
		SELECT * 
		FROM users 
		WHERE id = ?`,
		user.ID)

	if query.Error != nil {
		return nil, query.Error
	}

	err := repo.DB.Create(&user).Scan(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
