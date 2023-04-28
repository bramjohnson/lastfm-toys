package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	API_KEY string
}

func GetConfig() Config {
	viper.SetConfigFile("config.yml")
	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return config
}
