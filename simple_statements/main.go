package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("20x"); err == nil {
		fmt.Println("No error, i is:", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("Only one argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer, Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
		fmt.Printf("%T\n", args[1])
	}
}
