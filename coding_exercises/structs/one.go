package main

import "fmt"

func main()  {
	type person struct {
		name string
		age int
		colors []string
	}

	me := person{name: "Sean", age: 40, colors: []string{"blue", "green"}}

	you := person{name: "Sally Rider", age: 75, colors: []string{"white", "black", "pink"}}

	fmt.Printf("My info: %v\n", me)
	fmt.Printf("Sally's stuff: %v\n", you)
}