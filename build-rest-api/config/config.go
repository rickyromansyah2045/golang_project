package config

import (
	"build-rest-api/structs"
	"fmt"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Connection database Error", err)
	}

	db.AutoMigrate(structs.Person{})
	return db
}
