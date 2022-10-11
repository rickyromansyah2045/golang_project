package main

import (
	"fmt"
	"math"
	"sync"
)

type shape interface {
	area() float64
	parimeter() float64
}

type rectangle struct {
	widht, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.widht
}

func (c circle) parimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) parimeter() float64 {
	return 2 * (r.height + r.widht)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s parimeter %v\n", t, s.parimeter())
}

func main() {
	// var c1 shape = circle{radius: 5}
	// var r1 shape = rectangle{widht: 3, height: 2}

	// print("Rectangle", c1)
	// print("Circle", r1)

	// fmt.Printf("Type of c1: %T\n", c1)
	// fmt.Printf("Type of r1: %T\n", r1)

	// fmt.Println("Circle area", c1.area())
	// fmt.Println("Circle parimeter", c1.parimeter())

	// fmt.Println("Rectangle area", r1.area())
	// fmt.Println("Rectangle parimeter", r1.parimeter())

	// Interface
	// var randomValues interface{}

	// _ = randomValues

	// randomValues = "Jalan Sudirman"

	// randomValues = 20

	// randomValues = true

	// // randomValues = []string{"Airel", "Nanda"}

	// fmt.Println("Random Values", randomValues)

	// var v interface{}

	// v = 20

	// // v = v * 9

	// if value, ok := v.(int); ok == true {
	// 	v = value * 9
	// }

	// rs := []interface{}{1, "Airel", true, 2, "Ananda", true}

	// rm := map[string]interface{}{
	// 	"Name":   "Airel",
	// 	"Status": true,
	// 	"Age":    23,
	// }

	// _, _ = rs, rm
	// fmt.Println(rs...)
	// fmt.Println(rm)

	// // Reflect
	// var number = 23
	// var reflectNumber = reflect.ValueOf(number)

	// fmt.Println("tipe", reflectNumber.Type())
	// fmt.Println("tipe", reflectNumber.Interface())

	// var s1 = &student{Name: "john wick", Grade: 2}
	// fmt.Println("nama : ", s1.Name)

	// var reflectValue = reflect.ValueOf(s1)
	// var method = reflectValue.MethodByName("SetName")

	// method.Call([]reflect.Value{
	// 	reflect.ValueOf("wick"),
	// })

	// fmt.Println("Nama : ", s1.Name)

	// Goroutines
	// fmt.Println("main function start")

	// go secondProsess(8)

	// firstProses(8)

	// fmt.Println("No. of Gourutines", runtime.NumGoroutine())

	// time.Sleep(time.Second * 2)

	// fmt.Println("Main executed End")

	fruits := []string{"Apple", "Manggo", "Durian", "Rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)

	wg.Done()
}

// func firstProses(index int) {
// 	fmt.Println("First Proses func started")

// 	for i := 1; i <= index; i++ {
// 		fmt.Println("i=", i)
// 	}

// 	fmt.Println("End Proses func End")
// }

// func secondProsess(index int) {

// 	fmt.Println("Second Proses Func Started")

// 	for j := 1; j <= index; j++ {
// 		fmt.Println("J=", j)
// 	}

// 	fmt.Println("Second Proses End")
// }

// func (s *student) SetName(name string) {
// 	s.Name = name
// }
