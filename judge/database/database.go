package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/database/model"
)

var NanoDB *gorm.DB

func init() {
	err := Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func Connect() error {
	db, err := gorm.Open(sqlite.Open("./NanoOJ.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open database")
		return err
	}
	err = db.AutoMigrate(&model.Submission{}, &model.Problem{}, &model.User{}, &model.Contest{}, &model.Notification{})
	if err != nil {
		log.Println("Failed to migrate tables")
		return err
	}
	NanoDB = db

	return nil
}

func Close() error {
	db, err := NanoDB.DB()
	if err != nil {
		log.Println("Failed to get database connection")
		return err
	}
	err = db.Close()
	if err != nil {
		log.Println("Failed to close database")
		return err
	}

	return nil
}
