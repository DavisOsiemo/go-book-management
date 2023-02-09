package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This file returns a variable called DB

var (
	db *gorm.DB
)

// Gorm allows you to talk to Postgres, Mysqlite, Mysql
func Connect() {
	d, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
