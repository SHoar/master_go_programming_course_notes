package main

import "fmt"

func main() {

	schools := []string{"Jackson", "Lehigh", "huntington"}
	for i := 0; i < 3; i++ {
		if i <= len(schools) {
			fmt.Println(schools[i])
		}
	}

	// go's 'while' loop
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// how to write an infinite loop in go
	// sum := 0
	// for {
	// 	sum++
	// }
}
