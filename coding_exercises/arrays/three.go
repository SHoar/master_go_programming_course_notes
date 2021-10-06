package main

import f "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}
	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	// myArray[3] = 10.10 // cannot assign to index 3 b/c only has len = 3 [0, 1, 2]
	myArray[2] = 10.10

	f.Println(myArray)
}
