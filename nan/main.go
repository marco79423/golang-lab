package main

import (
	"fmt"
	"math"
)

func Test1() {
	a, b := 2, 3
	z := 0
	x, y := a/z, b/z
	fmt.Println(x, y, x == y) // True or False ?
}

func Test2() {
	a, b := 2.0, 3.0
	z := 0.0
	x, y := a/z, b/z
	fmt.Println(x, y, x == y) // True or False ?
}

func Test3() {
	a, b := math.NaN(), math.Sqrt(-1.0)
	fmt.Println(a, b, a == b) // True or False ?
}

func Test4() {
	a, b := math.NaN(), math.NaN()
	fmt.Println(a, b, a == b) // True or False ?
}

func Test5() {
	a, b := math.NaN(), math.NaN()
	fmt.Println("==", a == b) // True or False ?
	fmt.Println("!=", a != b) // True or False ?
	fmt.Println(">", a > b)   // True or False ?
	fmt.Println("<", a < b)   // True or False ?
}

func Test6() {
	m := make(map[float64]string)
	m[math.NaN()] = "NaN"
	m[math.NaN()] = "NaN"
	m[math.NaN()] = "NaN"
	fmt.Println(len(m), m)
}

func Test7() {
	m := make(map[float64]string)
	m[math.NaN()] = "NaN"
	m[math.NaN()] = "NaN"
	m[math.NaN()] = "NaN"

	fmt.Println(len(m), m)
	for k := range m {
		delete(m, k)
	}
	fmt.Println(len(m), m)
}

func main() {
	// Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
	Test7()
}
