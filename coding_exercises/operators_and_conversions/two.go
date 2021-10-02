package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// i to string (int to string)
	// iStr := strconv.Itoa(i) // int to ascii using strconv package
	iStr := fmt.Sprint(i)

	fmt.Printf("i is of type %T, with a value of %s\n", iStr, iStr)
	// s2 to int (string to int)
	s2Int, err := strconv.Atoi(s2)
	_ = err
	fmt.Printf("s2 is now a %T with a value of %d\n", s2Int, s2Int)

	// f to string (float64 to string)
	fStr := fmt.Sprint(f)
	fmt.Printf("f is now a %T with the value %s \n", fStr, fStr)

	// s1 to float64 (string to float64)
	s1F, err := strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Printf("s2 is now a %T with the value of %f\n", s1F, s1F)
}
