package main

import (
	f "fmt"
	"math"
)

func f1() {
	f.Println("This is f1() func")
}

func f2(a int, b int) {
	f.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	f.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 { // specifying float return
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) { // if specifying return value, need return keyword
	f.Println(s) // init at 0
	s = a + b
	return
}

// func defArgs(a int = 3, b int = 1){}  // no default arguments in Go

func main() {
	f1()
	f2(5, 7)
	f3(1, 2, 3, 4.015, 5.5, "six")
	p := f4(5.1)
	f.Println(p)
	a, b := f5(8, 9)
	f.Println(a, b)
	mySum := sum(5, 7)
	f.Println(mySum)
}
