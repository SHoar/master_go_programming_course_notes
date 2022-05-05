package main

import (
	"fmt"
	"strings"
)

func searchItem(a []string, b string) bool {
	truth := false
	for _, value := range a {
		result := strings.Compare(strings.ToLower(value), strings.ToLower(b))
		if result == 0 {
			truth = true
		}
	}
	return truth
}

func main() {
	// case sensitive search
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

	// case-insensitive search, added ToLower to value & b during compare
	animals = []string{"Lion", "tiger", "bear"}
	result = searchItem(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItem(animals, "lion")
	fmt.Println(result) // => true
}
