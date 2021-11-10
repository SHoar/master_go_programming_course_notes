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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
func main() {
	var s shape = circle{name: "circle", radius: 2.5}
	f.Printf("%T\n", s)

	// s.volume() -- doesn't have access in interface to volume
	// type assertion
	ball, ok := s.(circle)

	if ok {
		f.Printf("Ball Volume: %v\n", ball.volume())
	}

	s = rectangle{name: "rect", height: 2, width: 3}
	switch value := s.(type) {
	case circle:
		f.Printf("%#v has circle type\n", value)
	case rectangle:
		f.Printf("%#v has rectangle type\n", value)
	default:
		break
	}
}
