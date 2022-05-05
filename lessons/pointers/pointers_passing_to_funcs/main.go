package main

import f "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

func changeVar(a int) int {
	a = 66
	return a
}

func main() {
	x := 8
	p := &x	
	f.Println("the value of x before calling change()", x)
	change(p)
	f.Println("value of x after calling change:", x)
	changeVar(x)
	f.Println("new value of x after calling changeVar", x)
}
