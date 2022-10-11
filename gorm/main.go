package main

import (
	"fmt"
	"gorm/database"
	"gorm/models"
)

func main() {
	database.StartDB()

	createUser("ricky@hm.com")

}

func createUser(email string) {

	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	dbc := db.Create(User)

	if dbc.Error != nil {
		fmt.Println("Ada error", dbc.Error)
		return
	}

	fmt.Println(dbc)
}
