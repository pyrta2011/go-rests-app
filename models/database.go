package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=musicstore password=ft89yrwq sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Track{})

	DB = db
}
