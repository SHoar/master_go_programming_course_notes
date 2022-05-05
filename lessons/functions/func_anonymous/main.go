package main

import "fmt"

// func increment(x int) int { // not anonymous
// 	x++
// 	return x
// }

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous message and function")

	a := increment(10)

	fmt.Printf("%T\n", a())
	fmt.Printf("%d\n", a())
	fmt.Printf("%v\n", a())
}
