package main

import (
	f "fmt"
	"os"
	"strconv"
)

func main() {
	clArgs := os.Args[1:]

	var sum int
	var product int = 1

	for _, arg := range clArgs {
		if len(clArgs) >= 2 && len(clArgs) <= 10 {
			val, err := strconv.Atoi(arg)
			if err != nil {
				f.Println("error:", err)
			}
			sum += val
			product *= val

		} else {
			f.Println("You must enter between 2 and 10 args")
			break
		}
	}

	f.Printf("Sum: %d\n", sum)
	f.Printf("Product: %d\n", product)
}
