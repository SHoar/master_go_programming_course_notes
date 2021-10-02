package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	// scores := fmt.Sprintf(score)

	fmt.Printf("x: %d, y: %f, z: %s \n", x, y, z) // 1
	fmt.Printf("score vals are: %#v\n", score)

	for i := 0; i < len([4]{x,y,z,score}); i++ {
		fmt.Printf(i + " is %v \n", i)
	}

	fmt.Printf("All ints: %d, scores: %d\n", x, score) // 3
	fmt.Printf("Type of y: %T, score: %T\n", y, score)

}
