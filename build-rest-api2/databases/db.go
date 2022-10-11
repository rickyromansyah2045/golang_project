package databases

import (
	"build-rest-api2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	var host = "localhost"
	var port = "5432"
	var username = "postgres"
	var password = "Noerhick_02"
	var dbname = "db-go-sql"

	var conn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(conn))

	if err != nil {
		fmt.Println("Error open connection db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(models.Person{})

	if err != nil {
		fmt.Println("Error on migration", err)

		return Database{}, err
	}

	return Database{
		db: db,
	}, nil

}

func (d Database) GetPerson() ([]models.Person, error) {
	dbg := d.db.Find(&[]models.Person{})

	rows, err := dbg.Rows()

	if err != nil {
		return nil, err
	}

	persons := make([]models.Person, 0)

	for rows.Next() {
		var person models.Person

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)

		if err != nil {
			continue
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (d Database) CreatePerson(person models.Person) (models.Person, error) {
	dbg := d.db.Create(&person)

	row := dbg.Row()

	var personResult models.Person

	err := row.Scan(&personResult.ID, &personResult.FirstName, &personResult.LastName)

	return personResult, err

}
