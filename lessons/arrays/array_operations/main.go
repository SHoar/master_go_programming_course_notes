package main

import (
	f "fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	f.Printf("%#v\n", numbers)

	numbers[0] = 7
	f.Printf("%#v\n", numbers)

	// numbers[5] = 100

	// iterating w/ range keyword
	for i, v := range numbers {
		f.Printf("index: %v, value: %v \n", i, v)
	}

	// iterating C/C++/Java style
	for i := 0; i < len(numbers); i++ {
		f.Printf("index: %v, value: %v \n", i, numbers[i])
	}

	f.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		f.Println("index:", i, "value:", numbers[i])
	}

	// multi-dimensional array [Array of Arrays]
	balances := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}
	f.Println(balances)

	for _, arr := range balances {
		for _, value := range arr {
			f.Printf("%d\n", value)
		}
	}

	// equal operator creates a copyof an array
	m := [3]int{1, 2, 3}
	n := m
	f.Println("n is equal to m:", m == n) // true
	m[1] = -1
	f.Println("n is equal to m:", m == n) // false as m[1] has changed, but not n

}
