package main

import f "fmt"

func main() {
	
	birthYear := 1981
	currentYear := 2021

	age := 0
	
life: //label for loop
	for year := birthYear; year <= currentYear; year++ {
		age++
		f.Printf("Year: %d, age: %d\n", year, age)
		if year == currentYear {
			break life
		}
	}


	f.Println("Outside of life label loop")
}