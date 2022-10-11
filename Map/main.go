package main

import "fmt"

func main() {

	myMap := map[string]string{
		"Ruby":       "Is Beautifil",
		"Go":         "IS Very Fast",
		"JavaScript": "Hype",
	}

	value, isAvailbe := myMap["Go"]
	fmt.Println(value)
	fmt.Println(isAvailbe)
}
