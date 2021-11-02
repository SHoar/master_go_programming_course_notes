package main

import "fmt"

func sum(n ...int) int {
	sum := 0
	for index, value := range n {
		sum += value
		_ = index
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
