package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type FPTConfig struct {
	BaseUrl      string `mapstructure:"base_url"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
}
type config struct {
	Storage string `mapstructure:"storage"`
}

var Config config

func ReadConfigs() {
	var err error

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	currentDirectory, err := os.Getwd()
	configPath := filepath.Join(currentDirectory, "config")
	if err != nil {
		panic(err.Error())
	}

	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&Config)

}
