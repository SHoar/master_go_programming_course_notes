package main

import "fmt"

func main() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	var sum int

	for i, num := range nums {
		if i == 0 || i == nums[6] || i == nums[7] {
			continue
		} else {
			fmt.Printf("%v ", num)
			sum += num
		}
	}
	fmt.Println("\nsum:", sum)
}
