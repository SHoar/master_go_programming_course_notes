package main

import f "fmt"

func main() {
	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry

	f.Println("new value of z from x/y:", z)
}