package main

import "fmt"

func main(){
	type person struct {
		name string
		age int
		colors []string
	}

	me := person{name: "Sean", age: 40, colors: []string{"blue", "green"}}

	// you := person{name: "Sally Rider", age: 75, colors: []string{"white", "black", "pink"}}

	me.name = "Andrei"

	favoriteColors := []string{"Red", "Yellow", "Bue"}
	me.colors = favoriteColors

	for index, color := range me.colors {
		fmt.Println("Color:", color)
		_ = index
	}
}