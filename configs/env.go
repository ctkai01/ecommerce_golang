package configs

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	DB string `mapStructure:"DB"`
}

var EnvConfigs * envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err:= viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}