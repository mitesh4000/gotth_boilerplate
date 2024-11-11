package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"templeTest/model"
)

var DB *gorm.DB

func InitializeDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	DB = db
	err = DB.AutoMigrate(&model.Item{})
	if err != nil {
		panic("failed to migrate database")
	}
}
