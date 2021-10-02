package main

import "fmt"

func main() {
	type duration int

	var hour duration
	fmt.Printf("hour value %v and type %T\n", hour, hour)

	hour = 3600
	fmt.Println(hour)
}
