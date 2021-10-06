package main

import "fmt"

func main() {
	const a float64 = 5.1

	const b = 6.7

	fmt.Println(a, b)

	// if consts are typed, you cannot mix types or infer type coersion
	// const x int = 5
	// const y float64 = 2.2 * x

	// when const aren't typed, variables can infer types.
	const c = 5
	const d = 2.2 * c
	fmt.Println(d)

}
