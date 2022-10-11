package main

import (
	"errors"
	"fmt"
)

func main() {

	/* Channel */
	// c := make(chan string)

	// go introduce("Ricky", c)

	// go introduce("Ridho", c)

	// go introduce("Rifki", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// ms3 := <-c
	// fmt.Println(ms3)

	// students := []string{"Airel", "Mailo", "Indah", "Namo"}

	// for _, v := range students {
	// 	go introduce(v, c)
	// }

	// for i := 1; i <= 4; i++ {
	// 	print(c)
	// }

	// close(c)

	// /* Defer Function*/
	// defer callDeferFunc()

	// fmt.Println("Gw Anggota DPR RI")

	/* Exit */
	// defer fmt.Println("Invoke with defer")
	// fmt.Println("Before Exiting")
	// os.Exit(1)

	// var number int
	// var err error

	// // fmt.Println(numbers)
	// // fmt.Println(err)

	// number, err = strconv.Atoi("123GH")

	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// number, err = strconv.Atoi("123")

	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// defer catchErr()

	// var password string

	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	// fmt.Println(err.Error())
	// 	// Panic
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println(valid)
	// }

	// Channel
	c := make(chan int)

	go greet(c, 5)
	go greet(c, 10)

	fmt.Println("Ini di atas")

	fmt.Println(<-c)
	fmt.Println(<-c)

}

func greet(c chan int, val int) {
	c <- val
}

// Recover
func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Aplikasi Running Perfect")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password has to have more than 4 character")
	}

	return "Valid Password", nil
}

func callDeferFunc() {
	deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}
