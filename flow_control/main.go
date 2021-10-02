package main

import "fmt"

func main() {
	price, inStock := 100, true

	// if price > 80 {
	// 	fmt.Println("Too Expensive")
	if price <= 100 && inStock == true {
		fmt.Println("Buy it!")
	} else if price == 100 {
		fmt.Println("On the edge")
	}

	// _ = inStock

}
