package main

import f "fmt"

func main() {

	name := "Andrei" // name stores value of "Andrei"
	f.Println(&name) // mem adrr of name
	x := 2
	ptr := &x //memory addr of x

	f.Printf("ptr is of type %T with a value of %v and address of%p\n", ptr, ptr, &ptr)
	f.Printf("address of x is %p\n", &x)

	var ptr1 *float64 // zero val is nil
	_ = ptr1

	p := new(int)

	x = 100
	p = &x

	f.Printf("p is of type %T with a value of %v\n", p, p)
	f.Printf("address of x is %p\n", &x)

	*p = 90 // assign x = 90
	f.Println(x, *p)
	f.Println("*p==x", *p == x)
}
