package main

import f "fmt"

func main() {
	// n1 := []int{10, 20, 30, 40, 50}
	// n2 := append(n1, 100)
	// n2[0] = 66

	n1 := []int{10, 20, 30, 40, 50}
	n2 := n1[:3]
	n2[0] = 66
	f.Println("++", n1)
	f.Println(n2)
}
