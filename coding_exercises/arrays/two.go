package main

import f "fmt"

func main() {
	nums := [...]int{30, -1, -6, 90, -6}

	// write a go program that counts the number of positive even numbers in the array

	count := 0
	for i := range nums {
		if nums[i] > 0 && nums[i]%2 == 0 {
			count++
			f.Println("index:", i, "number:", nums[i]) // nums is assigned in ln 6
		}
	}
	f.Println("Done")
}
