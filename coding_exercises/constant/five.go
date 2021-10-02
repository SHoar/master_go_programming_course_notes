package main

import "math"

func main() {
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b

	x := 8
	const xc int = 8 // x cannot be assigned to a const b/c is a variable and can change

	var noIPv6 = math.Pow(2, 128)

	_, _ = x, noIPv6
}
