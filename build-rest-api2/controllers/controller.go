package controllers

import (
	"build-rest-api2/databases"
	"build-rest-api2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db databases.Database
}

func New(db databases.Database) Controller {
	return Controller{
		db: db,
	}
}

func (c Controller) GetPerson(ctx *gin.Context) {
	persons, err := c.db.GetPerson()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error get data",
		})
	}

	ctx.JSON(http.StatusOK, persons)
}

func (c Controller) CreatePerson(ctx *gin.Context) {
	var newPerson models.Person

	err := ctx.BindJSON(&newPerson)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error bind json Request",
		})
	}

	personResult, err := c.db.CreatePerson(newPerson)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error Create Person",
		})
	}

	ctx.JSON(http.StatusOK, personResult)
}

func (c Controller) UpdatePerson(ctx *gin.Context) {
	// id := ctx.Query("id")

	// firstName := ctx.PostForm("first_name")
	// lastName := ctx.PostForm("last_name")

	// var newPerson models.Person

	// err := ctx.BindJSON(&newPerson)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"code":    "500",
	// 		"message": "error Bin json Request",
	// 	})
	// }
}
