package main

import f "fmt"

func main() {
	var empty interface{}

	empty = 5
	f.Printf("%T\n", empty)

	empty = 79.99
	f.Printf("%T\n", empty)

	empty = []int{5, 10, 15}
	f.Printf("%T\n", empty)

	empty = append(empty.([]int), 25) // to use implemented type while empty interface, use type assertion

	f.Println(empty)
}
