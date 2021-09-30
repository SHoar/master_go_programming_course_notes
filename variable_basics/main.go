package main

import "fmt"

func main() {
	var age int = 30

	fmt.Println("Age:", age)

	var name = "Dan"
	// the _ is a blank identifier & can be used to either give variables blank names
	// _ = name
	// fmt.Println("Your name is:", _)

	s := name
	fmt.Print(s)

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)
	// OR to hide the variables that you don't plan on using from the runtime compiler
	_, _ = salary, firstName
	var i, j int = 5, 8
	j, i = i, j

	sum := 5 + 2.3
	fmt.Println(sum)
}
