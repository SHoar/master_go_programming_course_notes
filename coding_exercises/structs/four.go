package main

import "fmt"

func main() {
	type grades struct {
		grade  int
		course string
	}

	type person struct {
		name   string
		age    int
		colors []string
		grades []grades
	}

	type colors []string

	me := person{
		name:   "Sean",
		age:    40,
		colors: []string{"blue", "yellow"},
		grades: []grades{
			{grade: 4, course: "Spanish"},
		},
	}

	fmt.Printf("My grades: %+v\n", me)
}
