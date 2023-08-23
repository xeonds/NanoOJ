package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string
		Port int
		Side string
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}

	Admin []struct {
		Username string
		Password string
	}

	Judger struct {
		Daemons []string
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

func DefaultConfig() *Config {
	return &Config{
		Server: struct {
			Host string
			Port int
			Side string
		}{
			Host: "	",
			Port: 8080,
			Side: "web-judge",
		},
		Database: struct {
			Host     string
			Port     int
			User     string
			Password string
			Name     string
		}{
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "root",
			Name:     "nano-oj",
		},
		Admin: []struct {
			Username string
			Password string
		}{
			{
				Username: "admin",
				Password: "admin",
			},
		}, Judger: struct {
			Daemons []string
		}{
			Daemons: []string{"localhost"},
		},
	}
}
