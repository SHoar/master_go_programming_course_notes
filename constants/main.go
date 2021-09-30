package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const b = 0

	// the runtime compiler will give an error saying "days not divisible by 0"
	// 	# command-line-arguments
	// constants/main.go:14:19: division by zero
	fmt.Println(days / b)
}
