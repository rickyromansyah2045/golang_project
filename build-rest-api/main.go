package main

import (
	"build-rest-api/config"
	"build-rest-api/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()

	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/person?:id", inDB.GetPerson)
	router.GET("/person", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person", inDB.DeletePerson)

	router.Run(":3000")
}
