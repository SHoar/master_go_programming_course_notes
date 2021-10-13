package main

import f "fmt"

func main() {
	// var m1 map[int]int uninitialized will return nil
	m1 := make(map[int]int) // make always initializes, so doesn't return nil
	if m1 == nil {
		f.Printf("T: %T, V: %v\n", m1, m1)
	} else {
		f.Println("You didn't do something right")
		f.Println(m1, len(m1))
	}

	m2 := map[int]string{101001: "Employee 1", 101002: "Employee 2"}
	m2[10] = "Abba"

	f.Printf("K: %#v, V: %#v\n", 101001, m2[101001])
	f.Printf("K: %#v, V: %#v\n", 3, m2[3]) // uninitialized key gets "" value
}
