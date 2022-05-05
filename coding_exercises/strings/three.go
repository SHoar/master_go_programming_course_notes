package main

import f "fmt"

func main() {
	s1 := "țară means country in Romanian"

	asciiArr := []string{}
	for i, char := range s1 {
		f.Printf("index %d --> rune: %c \n", i, char)
		asciiArr = append(asciiArr, f.Sprintf("%#v", char))
	}

	f.Printf("Bytes in string: ")
	for l := 0; l < len(s1); l++ {
		f.Printf("%v ", s1[l])
	}
}
