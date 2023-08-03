package util

import (
	"happy-admin-service/helper"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HappyBaseUrlService     string `mapstructure:"HAPPY_BASE_URL_SERVICE"`
	HappyUserServiceUrl     string `mapstructure:"HAPPY_USER_SERVICE_URL"`
	HappyProductServiceUrl  string `mapstructure:"HAPPY_PRODUCT_SERVICE_URL"`
	HappyFeedbackServiceUrl string `mapstructure:"HAPPY_FEEDBACK_SERVICE_URL"`
	HappyOrderServiceUrl    string `mapstructure:"HAPPY_ORDER_SERVICE_URL"`
	AllowOrigins            string `mapstructure:"ALLOW_ORIGINS"`
	ServerAddr              string `mapstructure:"SERVER_ADDR"`
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
		happyUserSvcUrl := os.Getenv("HAPPY_USER_SERVICE_URL")
		happyProductSvcUrl := os.Getenv("HAPPY_PRODUCT_SERVICE_URL")
		happyFeedbackSvcUrl := os.Getenv("HAPPY_FEEDBACK_SERVICE_URL")
		happyOrderSvcUrl := os.Getenv("HAPPY_ORDER_SERVICE_URL")
		allowOrigins := os.Getenv("ALLOW_ORIGINS")
		serverAddr := os.Getenv("SERVER_ADDR")

		config := Config{
			HappyUserServiceUrl:     happyUserSvcUrl,
			HappyProductServiceUrl:  happyProductSvcUrl,
			HappyFeedbackServiceUrl: happyFeedbackSvcUrl,
			HappyOrderServiceUrl:    happyOrderSvcUrl,
			AllowOrigins:            allowOrigins,
			ServerAddr:              serverAddr,
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
