package main

import f "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	// mySlice = append(mySlice, 6)
	mySlice[1] = 6

	// a := 10 // quick convert to float
	a := 10.
	mySlice[0] = a

	// mySlice[3] = 10.10
	f.Println(len(mySlice), cap(mySlice))

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	f.Println(mySlice)
}
