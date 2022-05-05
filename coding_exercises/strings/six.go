package main

import f "fmt"

func main() {
	s := "你好 Go!"
	r := []rune(s)
	f.Println(r)

	for i, symbol := range r {
		f.Printf("Index: %d, rune: %c\n", i, symbol)
	}
}
