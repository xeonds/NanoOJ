package lib

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// 配置管理
func LoadConfig[Config any]() *Config {
	if _, err := os.Stat("config.yaml"); err != nil {
		confTmpl := new(Config)
		data, _ := yaml.Marshal(confTmpl)
		os.WriteFile("config.yaml", []byte(data), 0644)
		log.Fatal(errors.New("config file not found, a template file has been created"))
	}
	if err := func() error {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		return viper.ReadInConfig()
	}(); err != nil {
		log.Fatal(errors.New("config file read failed"))
	}
	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(errors.New("config file parse failed"))
	}
	return config
}
