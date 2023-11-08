package di

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/config"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/api"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/db"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases"
)

func InitializeAPI(c *config.DbConfig) (*api.ServerHttp, error) {

	gormDB, err := db.ConnectPostgres(c)
	if err != nil {
		return nil, err
	}

	adminRepo := repository.NewAdminRepository(gormDB)
	categoryRepo := repository.NewCategoryRepository(gormDB)
	dishRepo := repository.NewDishRepository(gormDB)
	sellerRepo := repository.NewSellerRepository(gormDB)
	userRepo := repository.NewUserRepository(gormDB)

	adminUcase := usecases.NewAdminUsecase(adminRepo, userRepo, sellerRepo, categoryRepo)
	sellerUcase := usecases.NewSellerUsecase(sellerRepo, dishRepo)
	userUcase := usecases.NewUserUsecase(userRepo, dishRepo, sellerRepo)

	sellerHandler := handlers.NewSellerHandler(sellerUcase)
	userHandler := handlers.NewUserHandler(userUcase)
	adminHandler := handlers.NewAdminHandler(adminUcase)

	serverHttp := api.NewServerHttp(adminHandler, sellerHandler, userHandler)

	return serverHttp, nil
}
