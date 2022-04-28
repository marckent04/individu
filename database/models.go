package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Birthday  string
	Tel string
}


func (u User) Create(db *gorm.DB) (tx *gorm.DB) {
	return db.Create(&u)
}

func (u User) Update(db *gorm.DB) (tx *gorm.DB) {
	return db.UpdateColumns(&u)
}

func (u User) Delete(db *gorm.DB) (tx *gorm.DB) {
	return db.Delete(&u)
}