package main

import (
	"fmt"
	"math"
	"strings"
)

type isOddNum func(int) bool

type Employee struct {
	name     string
	age      int
	division string
}

func main() {

	// Conditional (Temporary Variabel)
	// var currentYear = 2022

	// if age := currentYear - 1998; age < 17 {
	// 	fmt.Println("Kamu belum boleh membuat kartu SIM")
	// } else {
	// 	fmt.Println("Kamu sudah boleh membuat kartu SIM")
	// }

	// // Switch
	// var score = 7

	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// // Switch With relational Operator
	// var score = 2

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not dad")
	// default:
	// 	{
	// 		fmt.Println("syudy harder")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// // Switch With relational Operator
	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not dad")
	// 	// Fallthrough Keyword
	// 	fallthrough
	// case score > 5:
	// 	fmt.Println("It is oke, but please study harder")
	// default:
	// 	{
	// 		fmt.Println("syudy harder")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// // Nested Conditional
	// var score = 15

	// if score > 7 {
	// 	switch score {
	// 	case 10:
	// 		fmt.Println("Perfect")
	// 	default:
	// 		fmt.Println("Nice")
	// 	}
	// } else {
	// 	if score == 5 {
	// 		fmt.Println("not Bad")
	// 	} else if score == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if score == 0 {
	// 			fmt.Println("try harder")
	// 		}
	// 	}
	// }

	// // Looping
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// // Second Way of Looping
	// var i = 0

	// for i < 3 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// // Looping (third way of looping)
	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++

	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// // Looping (break and continoe keyboards)
	// for i := 1; i <= 20; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 18 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// // Nested Looping
	// for i := 1; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print("*")
	// 	}

	// 	fmt.Println()
	// }

	// 	// Looping label
	// outerloop:
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println("Perulangan ke -", i+1)
	// 		for j := 0; j < 3; j++ {
	// 			if i == 2 {
	// 				break outerloop
	// 			}
	// 			fmt.Print(j, " ")
	// 		}
	// 		fmt.Print("\n")
	// 	}

	// // Dasar Array Golang
	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Airell", "Nanda", "Mailo"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// // Array Modefy Element Through Index
	// var fruits = [3]string{"Apel", "Mangga", "Pisang"}
	// fruits[0] = "Apple"
	// fruits[1] = "Banana"
	// fruits[2] = "Watermelon"

	// fmt.Printf("%#v\n", fruits)

	// // Array Loop Through Element
	// var fruits = [4]string{"Apel", "Mangga", "Watermelon", "Tomat"}
	// // Cara Pertama
	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value: : %s\n", i, v)
	// }

	// fmt.Sprintln(strings.Repeat("#", 25))

	// // Cara Kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }

	// // Array Multidimensional Array
	// balances := [3][3]int{{5, 6, 7}, {8, 9, 10}, {3, 4, 5}}

	// for _, arr := range balances {
	// 	for _, value := range arr {
	// 		fmt.Printf("%d ", value)
	// 	}
	// 	fmt.Println()
	// }

	// // Slice
	// var fruits = []string{"Apple", "Banana", "Mango"}

	// // _ = fruits
	// // fmt.Printf("%#v\n", fruits)
	// fmt.Println(fruits)

	// // Slice Make Function
	// var fruits = make([]string, 6)

	// _ = fruits

	// fmt.Printf("%#v", fruits)

	// // Slice Append Function
	// var fruits = make([]string, 3)

	// // fruits[0] = "Apple"
	// // fruits[1] = "Banana"
	// // fruits[2] = "Mango"

	// fruits = append(fruits, "Banana", "Orange", "Apel")

	// fmt.Printf("%#v", fruits)

	// // Slice Append function with ellipsis
	// var fruits1 = []string{"Apple", "Banana", "Mango"}

	// var fruits2 = []string{"Durian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	// Function
	// greet("Ricky", "Jl Jendral Sudirman")

	// // Function Return
	// var names = []string{"Airel", "Jordan"}
	// var printMsg = greet("Hei", names)

	// fmt.Println(printMsg)

	// // Function (Returning multiple values)
	// var diameter float64 = 15

	// var area, circumference = calculate(diameter)

	// fmt.Println("Area:", area)
	// fmt.Println("Circumference:", circumference)

	// // Function (Predefined return value)
	// var diameter float64 = 15

	// var area, circumference = calculates(diameter)

	// fmt.Println("Area:", area)
	// fmt.Println("Circumference:", circumference)

	// Function (variadic function #1)
	// studentList := print("Airel", "Nanda", "Ricky", "Roman", "Adam")

	// fmt.Printf("%v", studentList)

	// // Function (Variadic Function #2)
	// numberList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// result := sum(numberList...)

	// fmt.Println("Result:", result)

	// Function (Variadic Function #3)
	// profile("Ricky", "pasta", "jagung bakar", "Hangga Kris", "Bobo koe")

	// Closure (Declare closure in variable)
	// var evenNumbers = func(numbers ...int) []int {
	// 	var result []int

	// 	for _, v := range numbers {
	// 		if v%2 == 0 {
	// 			result = append(result, v)
	// 		}
	// 	}

	// 	return result
	// }

	// var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	// fmt.Println(evenNumbers(numbers...))

	// Closure (IIFE)
	// var isPalindrome = func(str string) bool {
	// 	var temp string = ""

	// 	for i := len(str) - 1; i >= 0; i-- {
	// 		temp += string(byte(str[i]))
	// 	}

	// 	return temp == str
	// }("katak")

	// fmt.Println(isPalindrome)

	// Closure (Callback)
	// var numbers = []int{2, 5, 8, 10, 3, 99, 23}

	// var find = findOddNumbers(numbers, func(number int) bool {
	// 	return number%2 != 0
	// })

	// fmt.Println("Total odd numbers:", find)

	// Pointer
	// var firstNumber *int
	// var secondNumber *int
	// _, _ = firstNumber, secondNumber

	// // Pointer Memory Address
	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber

	// fmt.Println("FirstNumber (value) :", firstNumber)
	// fmt.Println("FirstNumber (memori address) :", &firstNumber)

	// fmt.Println("SecondNumber (value) : ", *secondNumber)
	// fmt.Println("SecondNumber (memori adddress) : ", secondNumber)

	// Struct
	// var employ Employee

	// employ.name = "Ridho"
	// employ.age = 22
	// employ.division = "Developer"

	// fmt.Println(employ.name)
	// fmt.Println(employ.age)
	// fmt.Println(employ.division)

	// Struct (Slice of anonymous struct)

	// Map Again
	// var person map[string]string

	// var people = []map[string]string{
	// 	{"name": "Airel", "age": "23"},
	// 	{"name": "Nanda", "age": "24"},
	// 	{"name": "Aurel", "age": "52"},
	// }

	// for i, person := range people {
	// 	fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	// }

	// // Aliase
	// type second = uint

	// var hour second = 3600
	// fmt.Printf("Hour Type: %T\n", hour)

	// Pointer
	// var a int = 10
	// fmt.Println("Before:", &a)

	// changeValue(&a)

	// fmt.Println("After", a)
}

func changeValue(number *int) {
	*number = 20
}

func findOddNumbers(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

func profile(name string, favFood ...string) {
	mergeFavFoods := strings.Join(favFood, ",")

	fmt.Println("Hello there !!! I'm", name)
	fmt.Println("I really lova eating", mergeFavFoods)
}

func sum(number ...int) int {
	total := 0

	for _, v := range number {
		total += v
	}
	return total
}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}

		result = append(result, temp)
	}

	return result
}

func calculates(d float64) (area float64, citcumference float64) {

	// Menghitung luas
	area = math.Pi * math.Pow(d/2, 2)
	// Menfhitung keliling

	citcumference = math.Pi * d

	return
}

func calculate(d float64) (float64, float64) {

	// Menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)
	// Menfhitung keliling

	var citcumference = math.Pi * d

	return area, citcumference
}

func greet(msg string, names []string) string {
	// fmt.Printf("Hello there! My name is %s and I'm %s years old", name, adress)

	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result

}
