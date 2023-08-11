package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){

	// get all credential from dotenv
	// reference for this tutorial: https://youtu.be/2UjXSFvecdU
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	// use .Sprintf to assign %s from variable we get from dotenv above
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
  	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})


	// sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbuser, dbpassword, dbhost, dbname))
  	// gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	

	if err != nil {
		panic("--> database failed to connect <--")
	}

	DB = db
	fmt.Println("--> successfully connected to database <--")

}