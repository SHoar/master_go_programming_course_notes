package main

import f "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} // you can slice an array
	// a[start:stop]
	b := a[1:3]
	f.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	f.Println(s2)

	f.Println(s1[2:]) // when stop index missing, do start index to finish
	// ^ same as s1[2:len(s1)]

	f.Println(s1[:3]) // same as s1[0:3]
	f.Println(s1[:])  // same as s1[0:len(s1)]
	// f.Println(s1[:100]) //s1 does not have up to index 100, error

	s1 = append(s1[:4], 100)
	f.Println(s1)

	s1 = append(s1[:4], 200) // able to overwrite existing values
	f.Println(s1)
	// a = append(s1, s2...) // cannot append to arrays
}
