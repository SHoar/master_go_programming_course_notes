package main

import f "fmt"

func main() {
	for x := 1; x <= 50; x++ {
		if x%7 == 0 {
			f.Println(x)
		}
	}
}
