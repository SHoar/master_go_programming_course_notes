package main

import (
	f "fmt"
)

type money float64

func (m money) print() {
	f.Printf("%.2f\n", m)
}

func main() {
	m := money(35.467)
	m.print()
}