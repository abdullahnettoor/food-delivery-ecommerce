// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/config"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/api"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/db"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases"
)

// Injectors from wire.go:

func InitializeAPI(c *config.DbConfig) (*api.ServerHttp, error) {

	// gormDB, err := db.ConnectDatabase(cfg)
	// if err != nil {
	// 	return nil, err
	// }
	
	// userRepository := repository.NewUserDataBase(gormDB)
	// userUseCase := usecase.NewuserUseCase(userRepository)
	// userHandler := handler.NewUserHandler(userUseCase)

	// adminRepository:= repository.NewAdminRepository(gormDB)
	// adminUseCase:= usecase.NewAdminUseCase(adminRepository)
	// adminHandler:=handler.NewAdminHandler(adminUseCase)


	// serverHTTP := http.NewServerHttp(userHandler,adminHandler)
	// return serverHTTP, nil

	gormDB, err := db.ConnectPostgres(c)
	if err != nil {
		return nil, err
	}

	adminRepo := repository.NewAdminRepository(gormDB)
	adminUcase := usecases.NewAdminUsecase(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminUcase)

	serverHttp := api.NewServerHttp(adminHandler)
	
	return serverHttp, nil
}