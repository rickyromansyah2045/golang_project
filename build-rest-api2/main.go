package main

import (
	"build-rest-api2/controllers"
	"build-rest-api2/databases"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.Start()

	if err != nil {
		fmt.Println("Error start databases", err)
	}

	ctl := controllers.New(db)

	r := gin.Default()

	r.GET("persons", ctl.GetPerson)
	r.GET("person/:id")
	r.POST("person", ctl.CreatePerson)
	r.PUT("/person/:id")
	r.DELETE("/person/:id")

	r.Run("localhost:9090")
}
