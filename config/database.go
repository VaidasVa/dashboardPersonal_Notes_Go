package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	//dsn := os.Getenv("DB_URL")
	dsn := "root:root@tcp(172.17.0.2:3306)/notes?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connecting to DB", err)
	}
	log.Println("Connected to DB")

}
