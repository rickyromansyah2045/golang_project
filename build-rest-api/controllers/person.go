package controllers

import (
	"build-rest-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get one data with (Id)
func (idb *InDB) GetPerson(c *gin.Context) {

	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All data
func (idb *InDB) GetPersons(c *gin.Context) {

	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create person
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	person.FirstName = firstName
	person.LastName = lastName
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

// Update data person
func (idb *InDB) UpdatePerson(c *gin.Context) {

	id := c.Query("id")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data Not Found",
		}
	}

	newPerson.FirstName = firstName
	newPerson.LastName = lastName

	err = idb.DB.Model(&person).Updates(newPerson).Error

	if err != nil {
		result = gin.H{
			"result": "Update Failed",
		}
	} else {
		result = gin.H{
			"result": "Update Success",
		}
	}

	c.JSON(http.StatusOK, result)
}

// Delete Data Person
func (idb *InDB) DeletePerson(c *gin.Context) {

	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "Delete Failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleter succefully",
		}
	}

	c.JSON(http.StatusOK, result)
}
