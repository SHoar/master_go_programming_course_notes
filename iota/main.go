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

	var array
	for var i int; i < 5; i++ {
		i = iota
	}

}
