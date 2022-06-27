package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// Connect()  connects with the database
func Connect() {
	d, err := gorm.Open("mysql", "akhil:demopassword/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns a gorm.DB
func GetDB() *gorm.DB {
	return db
}
