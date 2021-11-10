package main

import (
	f "fmt"
	"math"
	"strings"
)

type shape interface {
	area() float64
	perimiter() float64
}

type circle struct {
	name   string
	radius float64
}

type rectangle struct {
	name   string
	height float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimiter() float64 {
	return 2 * math.Pi * c.radius
}

// func (c circle) printCircle() {
// 	f.Println("Shape:", c)
// 	f.Println("Area:", c.area())
// 	f.Println("Perimeter:", c.perimiter())
// }

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimiter() float64 {
	return 2 * (r.height + r.width)
}

// func (r rectangle) printRectangle() {
// 	f.Println("Shape:", r)
// 	f.Println("Area:", r.area())
// 	f.Println("Perimeter:", r.perimiter())
// }

func print(s shape) {
	f.Printf("Shape: %+v\n", s)
	f.Println("Area:", s.area())
	f.Println("Perimeter:", s.perimiter())
	f.Println(strings.Repeat("*", 5))
}
func main() {
	var s shape
	f.Printf("%T\n", s)

	ball := circle{name: "circle", radius: 2.5}
	s = ball

	print(s)

	room := rectangle{name: "rect", height: 3, width: 2}
	s = room
	f.Printf("Type of s: %T\n", s)

}
