package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Settings struct {
	SecretKey string `mapstructure:"secret_key"`
	Database  struct {
		Uri      string `mapstructure:"uri"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

var settings *Settings

func InitSettings() {
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		log.Fatal("Failed to unmarshal config file: ", err)
	}
}

func GetSettings() *Settings {
	return settings
}
