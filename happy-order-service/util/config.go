package util

import (
	"happy-order-service/helper"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
}

func setupConfigFile(path, configName string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	return err
}

func LoadConfig(path string) Config {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "prod" {
		dbDriver := os.Getenv("DB_DRIVER")
		dbSource := os.Getenv("DB_SOURCE")
		serverAddr := os.Getenv("SERVER_ADDR")

		config := Config{
			DBDriver:   dbDriver,
			DBSource:   dbSource,
			ServerAddr: serverAddr,
		}

		return config
	}

	err := setupConfigFile(path, "dev")
	helper.DoPanicIfError(err)

	config := Config{}
	err = viper.Unmarshal(&config)
	helper.DoPanicIfError(err)

	return config
}
