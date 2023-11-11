package config

import "github.com/spf13/viper"

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
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&dbConfig); err != nil {
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
