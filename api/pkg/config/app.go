package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db         *gorm.DB
	dbName     = os.Getenv("MYSQL_DB_NAME")
	dbPort     = os.Getenv("MYSQL_DB_PORT")
	dbHost     = os.Getenv("MYSQL_DB_HOST")
	dbUsername = os.Getenv("MYSQL_DB_USERNAME")
	dbPassword = os.Getenv("MYSQL_DB_PASSWORD")
)

func Connect() {
	log.Print(os.Getenv("FOO"))

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
