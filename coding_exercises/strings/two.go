package main

import (
	f "fmt"
)

func main() {
	r := 'ă'
	f.Printf("%T\n", r)

	m1 := "ma"
	m2 := "m"

	mama := f.Sprintf("%s%s%c", m1, m2, r) // converted into fmt string
	f.Println(mama)

	mama = m1 + m2 + string(r) // concatenated way
	f.Println(mama)

}
