package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		json.NewEncoder(w).Encode(Employment)
		return

		// Kalau mau Output HTML
		// tpl, err := template.ParseFiles("template.html")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// tpl.Execute(w, Employment)
		// return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(Employment) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		Employment = append(Employment, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
