package app

import (
	"log"

	"golang-gorm-item-order/model/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	dsn := "root:ElxIvs2c@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	DB = db

	db.AutoMigrate(&domain.Order{}, &domain.Items{})
}

func GetDB() *gorm.DB {
	return DB
}
