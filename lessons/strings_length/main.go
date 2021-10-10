package main

import (
	f "fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	f.Println(len(s1))

	name := "ţară"
	f.Println(len(name))

	f.Println(utf8.RuneCountInString(name))
	f.Printf("Rune: %#v --> RuneCount %d --> len: %d\n", "ă", utf8.RuneCountInString("ă"), len("ă")) //%c for byte
	f.Println(utf8.RuneCountInString("ă"))
}
