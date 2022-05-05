package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}
	fmt.Println("m: ", m)
	delete(m, 3)
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}
}
