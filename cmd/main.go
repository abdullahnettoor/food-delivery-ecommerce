package main

import (
	"log"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/config"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/di"
)

func main() {
	dbCfg, configErr := config.LoadDbConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	cldCfg := config.LoadImageUploader()

	server, diErr := di.InitializeAPI(dbCfg, cldCfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	}

	server.Start()

}
