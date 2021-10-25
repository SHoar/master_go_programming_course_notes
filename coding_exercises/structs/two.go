package main

import "fmt"

func main() {
	type person struct {
		name string
		age int
		colors []string
	}

	me := person{name: "Sean", age: 40, colors: []string{"blue", "green"}}

	you := person{name: "Sally Rider", age: 75, colors: []string{"white", "black", "pink"}}

	me.name = "Andrei"

	favoriteColors := []string{"Red", "Yellow", "Bue"}
	fmt.Printf("favorite colors: %T, %+v\n", favoriteColors, favoriteColors)
	me.colors = favoriteColors
	fmt.Printf("My favorite colors: %v\n", me.colors)

	you.colors = append(you.colors, "purple")

	fmt.Printf("Your favorite colors: %+v\n", you.colors)

	fmt.Printf("Me: %v\n", me)
	fmt.Printf("You: %v\n", you)

}