package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	// convert i to float64
	i_converted := float64(i)
	fmt.Printf("i is now %T with a value of %v\n", i_converted, i_converted)

	// convert f to int
	f_int := int(f)
	fmt.Printf("f is now %T with a value of %v\n", f_int, f_int)
}
