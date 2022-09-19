package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type DB struct {
	DSN string `mapstructure:"dsn"`
}
type config struct {
	Storage string `mapstructure:"storage"`
	DB      DB     `mapstructure:"db"`
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
