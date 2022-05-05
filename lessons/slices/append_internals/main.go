package main

import f "fmt"

func main() {
	var nums []int
	f.Printf("%#v\n, length: %d, capacity: %d\n", nums, len(nums), cap(nums))

	nums = append(nums, 1, 2)
	f.Printf("%#v\n, length: %d, capacity: %d\n", nums, len(nums), cap(nums))

	nums = append(nums, 3)
	f.Printf("%#v\n, length: %d, capacity: %d\n", nums, len(nums), cap(nums))

	nums = append(nums, 4)
	f.Printf("%#v\n, length: %d, capacity: %d\n", nums, len(nums), cap(nums))

	nums = append(nums, 100)
	f.Printf("%#v\n, length: %d, capacity: %d\n", nums, len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	f.Printf("%#v\n, length: %d, capacity: %d\n", nums, len(nums), cap(nums))

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")

	f.Println(letters, len(letters), cap(letters))
	// f.Println(letters[4])
	f.Println(letters[3:6])
}
