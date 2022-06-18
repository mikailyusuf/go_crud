package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_NAME = "users"
const DB_USER_NAME = "root"
const DB_PASSWORD = "root"
const DB_HOST = "127.0.0.1"
const DB_PORT = "8889"

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = connectTODB()
	return DB

}

func connectTODB() *gorm.DB {
	var err error
	dsn := DB_USER_NAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
	}
	return db
}
