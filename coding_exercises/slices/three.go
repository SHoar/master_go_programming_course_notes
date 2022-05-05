package main

import f "fmt"

func main() {
	nums := []float64{2.1, 3.2, 5.4}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.1, 6.6)
	f.Println(nums)
	n := []float64{1.1, 2.2}
	nums = append(nums, n...)
	f.Println(nums)
}
