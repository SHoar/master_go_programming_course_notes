package main

import "fmt"

func main() {
	const y = 2
	const x = 2.1
	fmt.Println(x * y)

	a := 5
	b := 2.1
	// fmt.Println(a*b) //mismatched types in and float64
	fmt.Println(a * int(b))
}
