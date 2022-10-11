package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func parsing_json() {

	var jsonString = `
		{
			"full_name": "Ricky Romansyah",
			"email": "rickyromansyah54@gmail.com",
			"age": 23
		}
	`

	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name :", result.FullName)
	fmt.Println("email :", result.Email)
	fmt.Println("age :", result.Age)
}

func json_map() {
	var jsonString = `
		{
			"full_name": "Ricky Romansyah",
			"email": "rickyromansyah54@gmail.com",
			"age": 23
		}
	`

	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("First Name", result["full_name"])
	fmt.Println("Email", result["email"])
	fmt.Println("Age", result["age"])
}

func Empty_Interface() {

	var jsonString = `
		{
			"full_name": "Ricky Romansyah",
			"email": "rickyromansyah54@gmail.com",
			"age": 23
		}
	`

	var temp interface{}

	var err = json.Unmarshal([]byte(jsonString), &temp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = temp.(map[string]interface{})

	fmt.Println("First Name", result["full_name"])
	fmt.Println("Email", result["email"])
	fmt.Println("Age", result["age"])

}

func Array_slice() {

	var jsonString = `[
		{
			"full_name": "Ricky Romansyah",
			"email": "rickyromansyah54@gmail.com",
			"age": 23
		},
		{
			"full_name": "Rifki Hikmawan",
			"email": "rifki@gmail.com",
			"age": 30
		}
	]
	`
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}

}

func main() {
	Array_slice()
}
