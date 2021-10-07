package main

import f "fmt"

func main() {
	var n []int
	f.Println(n == nil) // not initialized

	m := []int{}
	f.Println(m == nil) //initialized to 0

	// var c []int
	// f.Println(n == c) // invalid operation: n == c (slice can only be compared to nil)

	// var n1 []int = []int{2, 4, 5}
	// var c1 []int = []int{2, 4, 5}
	//f.Println(n1 == c1) // invalid operation: n == c (slice can only be compared to nil)
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true

	a = nil
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		f.Println("a and b slices are equal")
	} else {
		f.Println("a and b slices are not equal")
	}
}
