package main

import f "fmt"

func main() {
	numbers := []int{2, 3}

	numbers = append(numbers, 10) // returns a new slice
	f.Println(numbers)

	numbers = append(numbers, 20, 30, 40)
	f.Println(numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...) // by adding ellipsis to n, adds just values to existing numbers slice
	f.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, 2)
	nn := copy(dst, src)
	f.Println(src, dst, nn)
}
