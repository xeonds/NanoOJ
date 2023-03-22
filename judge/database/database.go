package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/model"
)

var NanoDB *gorm.DB

func init() {
	Connect()
}

func Connect() error {
	var err error
	NanoDB, err = gorm.Open(sqlite.Open("./NanoOJ.db"), &gorm.Config{})
	if err != nil {
		log.Println("Open database failed")
		return err
	}
	NanoDB.AutoMigrate(&model.Submission{})
	NanoDB.AutoMigrate(&model.Problem{})
	NanoDB.AutoMigrate(&model.User{})

	return nil
}

func Close() error {
	db, err := NanoDB.DB()
	if err != nil {
		log.Println("Get database connection failed")
		return err
	}
	err = db.Close()
	if err != nil {
		log.Println("Close database failed")
		return err
	}
	return nil
}
