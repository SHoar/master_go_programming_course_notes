package main

import (
	f "fmt"
)

type money float64

func (m money) print() {
	f.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return f.Sprintf("%.2f", m)
}

func main() {
	m := money(35.467)
	m.print()

	n := money(45.789678)
	f.Println(n.printStr())
}