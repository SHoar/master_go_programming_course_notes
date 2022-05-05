package main

import "fmt"

func main() {
	// iota is an iterative for consts
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		s3 = iota + 5
		s4
		s5
	)
	fmt.Println(s3, s4, s5)

}
