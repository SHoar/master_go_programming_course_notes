package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	a = int(b)

	fmt.Println(a, b)

	// var x int
	// x = "5"

	var value int
	var prince float64
	var name string
	var done bool
	// by default, an instantiated and typed var will get a 0 value for nums
	// "" for strings and 'false' for bools
	fmt.Println(value, prince, name, done)
}
