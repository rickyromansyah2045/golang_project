package main

import "fmt"

func main() {
	// fmt.Println("Hello World")

	// Cara Penulisan Variabel Pertama
	// var name string = "Ricky"
	// var age int = 22

	// fmt.Println(name)
	// fmt.Println(age)

	// // Cara Kedua
	// var nameTwo string
	// var ageTwo int

	// nameTwo = "Ridho"
	// ageTwo = 25

	// fmt.Println(nameTwo)
	// fmt.Println(ageTwo)

	// // Cara ketiga
	// var nameThree = "Galis"
	// var ageThree = 30

	// fmt.Printf("%T, %T\n", nameThree, ageThree)

	// // Cara Ke Empat
	// nameFour := "Aji"
	// ageFour := 40

	// fmt.Println(nameFour)
	// fmt.Println(ageFour)

	// // Multiple Variabel 1 tipe data
	// var student1, student2, student3 string = "satu", "dua", "tiga"
	// var first, second, third int

	// first, second, third = 1, 2, 3

	// fmt.Println(student1, student2, student3)
	// fmt.Println(first, second, third)

	// // Multiple Variabel Beda tipe data
	// var name, age, address = "Hanida", 27, "Jl Cempaka Putih"

	// one, two, three := "1", 2, "3"

	// fmt.Println(name, age, address)
	// fmt.Println(one, two, three)

	// // Variabel yang tidak perlu di pakai ya Guys
	// var firstVariabel string
	// var nama, tahun, alamat = "Siti", 2022, "Jalan Jendral Sudirman"

	// _, _, _, _ = firstVariabel, nama, tahun, alamat

	// // Verb Fmt
	// var name = "Aifal"

	// var age = 22

	// var address = "Jl Sudirman"

	// fmt.Printf("Hallo namaku %s, Umur aku adalah %d, dan aku tinggal di %s", name, age, address)

	// // Data Types
	// // Tipe Data Integer
	// var first uint8 = 89
	// var second int8 = -12

	// fmt.Printf("tipe data first : %T\n", first)
	// fmt.Printf("tipe data second : %T\n", second)

	// // Tipe Data Float
	// var decimalNumber float32 = 3.69

	// fmt.Printf("Decimal Number: %f\n", decimalNumber)
	// fmt.Printf("Decimal Number: %.3f\n", decimalNumber)

	// // Tipe Data Bool
	// var conditional bool = true

	// fmt.Printf("Ini tipe data bool? %t \n", conditional)

	// // String Biasa
	// var stringBiasa string = "Hallo"

	// fmt.Print(stringBiasa)

	// // String Tidak Escape
	// var message string = `Selamat Datang Ricky;) "Hahaha"`

	// fmt.Println(message)

	// Constanta
	// const name string = "Roman"

	// fmt.Println(name)

	// const fullName string = "Ricky Romansyah"

	// fmt.Println(fullName)

	// // Operator
	// var value = (2 + 2) * 3
	// fmt.Println(value)

	// var firstConditional = 2 < 3
	// var secondConditional = "ricky" == "Ricky"
	// var thirdConditional = 18 == 2.3
	// var fourthConditional = 11 <= 11

	// fmt.Println("First Conditional", firstConditional)
	// fmt.Println("Second Conditional", secondConditional)
	// fmt.Println("Third Conditional", thirdConditional)
	// fmt.Println("Fouth Conditional", fourthConditional)

	// Operator Logika (Logical Operator)
	var right = true
	var wrong = false

	var WrongAndRight = right && wrong
	var WrongOrRight = right || wrong
	var WrongSameRight = !wrong

	fmt.Printf("Wrong & Right %t\n", WrongAndRight)
	fmt.Printf("Wrong OR Right %t\n", WrongOrRight)
	fmt.Printf("Wrong Same Right %t\n", WrongSameRight)
}
