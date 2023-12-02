package config

import (
	"bytes"
	"log"

	embedfiles "github.com/abdullahnettoor/food-delivery-eCommerce"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string `mapstructure:"DBHOST"`
	User     string `mapstructure:"DBUSER"`
	Password string `mapstructure:"DBPASSWORD"`
	Name     string `mapstructure:"DBNAME"`
	Port     string `mapstructure:"DBPORT"`
	TimeZone string `mapstructure:"DBTIMEZONE"`
}

func LoadDbConfig() (*DbConfig, error) {
	var dbConfig DbConfig

	viper.AddConfigPath("/")

	env, err := embedfiles.ENV.ReadFile(".env")
	if err != nil {
		return nil, err
	}

	viper.SetConfigType("env")

	if err := viper.ReadConfig(bytes.NewBuffer(env)); err != nil {
		log.Println(err)
		return nil, err
	}

	if err := viper.Unmarshal(&dbConfig); err != nil {
		log.Println(err)
		return nil, err
	}

	return &dbConfig, nil
}

type ImgUploaderCfg struct {
	CloudUrl string
}

func LoadImageUploader() *ImgUploaderCfg {
	var cfg ImgUploaderCfg
	cfg.CloudUrl = viper.GetString("IMG_CLOUD_URL")
	return &cfg
}
