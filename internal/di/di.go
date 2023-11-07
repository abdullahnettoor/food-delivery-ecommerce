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
	adminUcase := usecases.NewAdminUsecase(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminUcase)

	sellerRepo := repository.NewSellerRepository(gormDB)
	sellerUcase := usecases.NewSellerUsecase(sellerRepo)
	sellerHandler := handlers.NewSellerHandler(sellerUcase)

	userRepo := repository.NewUserRepository(gormDB)
	userUcase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUcase)

	serverHttp := api.NewServerHttp(adminHandler, sellerHandler, userHandler)

	return serverHttp, nil
}
