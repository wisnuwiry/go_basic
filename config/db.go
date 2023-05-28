package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-trial-class/entity"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=rahasia123 dbname=trial_class port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	} else {
		fmt.Println("DB connected")
		DB = db
	}

	db.AutoMigrate(&entity.User{})
}
