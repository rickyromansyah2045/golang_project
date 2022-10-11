package main

import (
	"fmt"
	"net/http"
	"project-six/handler"
)

var PORT = ":8080"
var TEXT = "Application is listening on port"
var HELLO = "Hello World"

func main() {

	http.HandleFunc("/", handler.Greet)

	http.HandleFunc("/employees", handler.GetEmployee)

	http.HandleFunc("/employee", handler.CreateEmployee)

	fmt.Println(TEXT, PORT)
	http.ListenAndServe(PORT, nil)
}
