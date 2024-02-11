package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Printf("Config file not found; using default values")
		}
	}
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Printf("Config file not found; using default values")
		}
	}

	var config Config
	err := viper.Unmarshal(&config)
	return &config, err
}
