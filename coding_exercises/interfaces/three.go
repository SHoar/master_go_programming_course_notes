package main

import f "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube)) // since volume is being called on an implemented emptyinterface, use type assertion
	f.Printf("Cube Volume: %v\n", v)	
}