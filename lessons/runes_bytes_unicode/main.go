package main

import (
	f "fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b'
	f.Printf("Type: %T, Val: %d\n", var1, var2)

	str := "tana"
	f.Println(len(str))

	f.Println("byte (not rune) at position1:", str[1]) // slicing a string will return a byte

	for i := 0; i < len(str); i++ {
		f.Printf("%c", str[i])
	}
	f.Println("\n" + strings.Repeat("#", 20))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		f.Printf("%c", r)
		i += size
	}

	f.Println("\n" + strings.Repeat("#", 20))

	for _, r := range str {
		f.Printf("%c", r)
	}
}
