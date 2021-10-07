package main

import (
	"fmt"
	"strconv"
)

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	// scores := fmt.Sprintf(score)

	fmt.Printf("x: %d, y: %f, z: %s \n", x, y, z) // 1
	fmt.Printf("score vals are: %#v\n", score)
	intZ, _ := strconv.Atoi(z)
	for i := 0; i < len([]int{x, int(y), intZ}); i++ { // does not work because slice must have all of same type
		fmt.Printf("i is %v \n", i)
	}

	fmt.Printf("All ints: %d, scores: %v\n", x, score) // 3
	fmt.Printf("Type of y: %T, score: %T\n", y, score)

}
