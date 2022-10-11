package database

import (
	"fmt"
	"gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "Noerhick_02"
	dbPort   = "5432"
	dbname   = "learning-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbName=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{}, models.People{})
}

func GetDB() *gorm.DB {
	return db
}

func CreateUser(user models.User) error {

	dbc := db.Create(user)

	if dbc.Error != nil {
		fmt.Println("Error create user", dbc.Error)
		return dbc.Error
	}

	rows, err := dbc.Rows()

	if err != nil {
		fmt.Println("Error get row", err)
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email)

		if err != nil {
			fmt.Println("Error scan row", err)
			continue
		}
	}

	return nil
}
