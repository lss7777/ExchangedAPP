package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: %v ", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal("Error unable to decode struct: %v ", err)
	}
	initDB()
}
