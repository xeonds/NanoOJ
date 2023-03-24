package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Judge struct {
		Endpoint string
		Token    string
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

