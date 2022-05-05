package main

import f "fmt"

func main() {
	s1 := "I love Golang!"
	f.Println(s1[2:5])

	s2 := "Dosvidanya"
	f.Println(s2[0:2])

	rs := []rune(s2) // a new backing array is created
	f.Printf("%T\n", rs)

	f.Println(string(rs[0:3]))

	t1 := "Go is cool!"
	f.Printf("%v, %#v, %T \n", t1[0], t1[0], t1[0])
}
