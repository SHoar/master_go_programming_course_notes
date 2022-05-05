package main

import f "fmt"

func main() {
	n1 := []string{"one", "two", "three"}
	for index, value := range n1 {
		f.Printf("index: %d, value: %s\n", index, value)
	}
}
