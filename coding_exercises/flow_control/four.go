package main

import f "fmt"

func main() {
	for x := 1; x <= 500; x++ {
		if x%7 == 0 && x%5 == 0 {
			f.Printf("%d\n", x)
		}
	}
}