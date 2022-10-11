package main

import "fmt"

func main() {
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "Playstation")
	gamingConsole = append(gamingConsole, "Nintendo")
	gamingConsole = append(gamingConsole, "Gembot")
	gamingConsole = append(gamingConsole, "X BOX")

	for _, console := range gamingConsole {
		fmt.Println(console)
	}

}
