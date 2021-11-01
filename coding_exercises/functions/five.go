package main

import "fmt"

func sum(n ...int) (x int) {
	for index, value := range n {
		x += value
		_ = index
	}
	return
}

func main() {
	fmt.Println(sum(1,2,3,4,5))

}