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

	// for and continue
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// for and break
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	// the break statement does not terminate execution, only the for loop
	fmt.Println("message after for loop")
}
