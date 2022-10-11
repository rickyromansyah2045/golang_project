package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "Noerhick_02"
	Dbname   = "db-go-sql"
)

var (
	Db  *sql.DB
	Err error
)

func Greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Selamat Datang"

	fmt.Fprint(w, msg)
}

func ConnDb() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	Db, Err = sql.Open("postgres", psqlInfo)

	if Err != nil {
		panic(Err)
	}

	// defer Db.Close()

	Err = Db.Ping()
	if Err != nil {
		panic(Err)
	}

	fmt.Println("Successfully Connected to database")
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {

		name := r.FormValue("full_name")
		email := r.FormValue("email")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		var employee = Employee{}

		sqlStatement := `
		INSERT INTO employees (full_name, email, age, division)
		VALUES ($1, $2, $3, $4)
		Returning *`

		Err = Db.QueryRow(sqlStatement, name, email, convertAge, division).
			Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if Err != nil {
			panic(Err)
		}

		json.NewEncoder(w).Encode(employee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		var results = []Employee{}

		sqlStatement := `SELECT * FROM employees`

		rows, err := Db.Query(sqlStatement)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var employee = Employee{}

			err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

			if err != nil {
				panic(err)
			}

			results = append(results, employee)
		}

		json.NewEncoder(w).Encode(results)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)

}

func UpdateEmployee() {

	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;`

	res, err := Db.Exec(sqlStatement, 1, "Roman Picisan", "rickyromansyah54@gmail.com", "Developer", 22)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Update Data Amount", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employees WHERE id = $1;
	`

	res, err := Db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount", count)

}
