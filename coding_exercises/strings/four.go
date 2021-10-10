package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x) // ascii byte code

	// s1[0] = 'x' // strings are immutable

	// printing the number of runes in the string
	// fmt.Println(len(s1)) // only the length of the string

	// var runeSlice []byte
	// for i := 0; i < len(s1); i++ {
	// 	runeSlice = append(runeSlice, s1[i])
	// 	fmt.Printf("Index: %d, rune: %c\n", i, s1[i])
	// }
	// fmt.Println(len(runeSlice))

	// quickest way to cocunt the runes
	fmt.Println(utf8.RuneCountInString(s1))
}
