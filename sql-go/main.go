package main

import (
	"fmt"
	"net/http"
	"sql-go/helpers"
)

var (
	PORT = ":8080"
	TEXT = "Application Running on PORT"
)

func main() {
	helpers.ConnDb()

	http.HandleFunc("/", helpers.Greet)

	http.HandleFunc("/employees", helpers.GetEmployees)

	http.HandleFunc("/employee", helpers.CreateEmployee)

	fmt.Println(TEXT, PORT)
	http.ListenAndServe(PORT, nil)

}
