package config

import (
	"ass-03/models"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "pratama21"
	dbPort   = "8082"
	dbname   = "db_ass3"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connecting to Database : ", err)
	}

	db.Debug().AutoMigrate(models.Weather{})
}

func GetDB() *gorm.DB {
	return db
}