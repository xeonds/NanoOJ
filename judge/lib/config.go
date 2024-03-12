package lib

import (
	"errors"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// 配置管理
func LoadConfig[Config any]() (*Config, error) {
	if _, err := os.Stat("config.yaml"); err != nil {
		confTmpl := new(Config)
		data, _ := yaml.Marshal(confTmpl)
		os.WriteFile("config.yaml", []byte(data), 0644)
		return nil, errors.New("config file not found")
	}
	if err := func() error {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		return viper.ReadInConfig()
	}(); err != nil {
		return nil, errors.New("config file read failed")
	}
	if err := viper.Unmarshal(new(Config)); err != nil {
		return nil, errors.New("config file parse failed")
	}

	return viper.Get("config").(*Config), nil
}
