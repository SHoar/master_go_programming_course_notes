package main

import f "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	f.Println(grades)

	accounts := [3]int{2: 50}
	f.Println(accounts)

	names := [...]string{5: "Dan"}
	f.Println(names, len(names))

	// skipping keyed index
	cities := [...]string{5: "Paris",
		"London", // gets index from last key element +1
		1:        "NYC"}
	f.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true} // bools arrays will assign default values as false
	f.Printf("%#v\n", weekend)
}
