package config

import (
	"Assignment-02/structs"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

func DbInit() *gorm.DB  {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&structs.Order{}, &structs.Item{}, )
	return db
}