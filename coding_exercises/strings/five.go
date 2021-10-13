package main

import (
	f "fmt"
)

func main() {
	s := "你好 Go!"

	// convert the string
	s2 := []byte(s)
	// 1. convert the string to a byte slice
	var byteSlice []byte
	for i := 0; i < len(s); i++ {
		byteSlice = append(byteSlice, s[i])
		f.Printf("index: %d --> rune: %c\n", i, s2[i])
	}

	// 2 print the byte slice
	f.Println(byteSlice)
	f.Println(s2)
}
