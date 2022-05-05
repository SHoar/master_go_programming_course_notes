package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1

	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)
	fmt.Printf("*p1 is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)

	// mind blown!
	fmt.Printf("**pp1 is %v\n", **pp1)
	**pp1++
	fmt.Printf("a is now %v\n", a)

	//**COMPARING POINTERS**//
	var p2 *int // nothing assigned, is nil
	fmt.Println(p2)
	fmt.Println(&p2) // but it has a mem addr

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) // false b/c they point to different variables

	p4 := &y
	fmt.Println(p2 == p4) // true b/c they point to the same variable
}