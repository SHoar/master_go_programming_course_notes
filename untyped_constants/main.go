package main

import "fmt"

func main() {
	const a float64 = 5.1

	const b = 6.7

	fmt.Println(a, b)

	// const x int = 5
	// const y float64 = 2.2 * x

	const c = 5
	const d = 2.2 * c
	fmt.Println(d)

}
