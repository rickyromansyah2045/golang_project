package main

import "fmt"

func main() {
	languages := [...]string{
		"Java",
		"Golang",
		"PHP",
		"Ruby",
		"JavaScript",
		"Kotlin",
	}

	fmt.Println(languages)
	lenght := len(languages)
	fmt.Println(lenght)

	for index, lang := range languages {
		fmt.Println("Index", index, "Language", lang)
	}

}
