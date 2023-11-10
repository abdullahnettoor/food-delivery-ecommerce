package repository

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
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
		return nil, e.ErrNotFound
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
		return nil, e.ErrNotFound
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
		WHERE email = ?
		OR phone = ?`,
		user.Email, user.Phone)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected != 0 {
		return nil, e.ErrConflict
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

func (repo *UserRepository) Verify(phone string) error {
	fmt.Println("Phone is", phone)
	err := repo.DB.Exec(`
	UPDATE users
	SET status = 'Active'
	WHERE phone = ?`, phone).Error

	if err != nil {
		return err
	}

	return nil
}
func (repo *UserRepository) Block(id string) error {
	err := repo.DB.Exec(`
	UPDATE users
	SET status = 'Blocked'
	WHERE id = ?`, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Unblock(id string) error {
	err := repo.DB.Exec(`
	UPDATE users
	SET status = 'Active'
	WHERE id = ?`, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) DeleteByPhone(phone string) error {
	err := repo.DB.Exec(`
	DELETE FROM users
	WHERE phone = ?`, phone).Error

	if err != nil {
		return err
	}

	return nil
}
