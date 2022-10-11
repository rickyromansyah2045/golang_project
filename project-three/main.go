package main

import (
	"fmt"
	"project-three/helpers"
)

func main() {
	fmt.Println("Hello")

	helpers.Greet()
	var person = helpers.Person{}

	person.Invokegreet()
}
