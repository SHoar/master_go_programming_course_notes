package main

import f "fmt"

func main() {
	//Every element in an array must be of same type, and the array has a fixed length
	var numbers [4]int
	f.Printf("%v\n", numbers)
	f.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	f.Printf("%v, a1\n", a1)

	var a2 = [3]int{-10, 1, 100}
	f.Printf("a2 %#v\n", a2)

	a3 := [3]string{"Diana", "Paul", "John"}
	f.Printf("a3 %#v\n", a3)

	a4 := [4]string{"x", "y"}
	f.Printf("a4 %#v\n", a4)

	a5 := [...]int{1, 2, 3, -1, 10, 60}
	f.Printf("a5 %#v\n", a5)
	f.Printf("The length of a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3,
		4,
		5,
	}
	f.Printf("a6 %v\n", a6)
}
