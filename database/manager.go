package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnect(src string)  (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open(src), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	return db, err
}

func LaunchMigrations(db *gorm.DB )  {
	db.AutoMigrate(&User{})
}