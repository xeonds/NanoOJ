package controller

import (
	"log"

	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/lib"
)

var db *gorm.DB
var repo database.Repository

func init() {
	config, err := lib.LoadConfig[config.Config]()
	if err != nil {
		log.Fatal("Failed to load config file")
	}
	db = lib.NewDB(&config.DatabaseConfig, nil)
	repo = database.Repository{DB: db}
}
