package main

import f "fmt"

func main() {
	count := 0
	for x:=1; x <= 50; x++ {
		if x%7 != 0 {
			continue
		}
		f.Printf("%d\n", x)
		count++
		if count == 3 {
			break
		}
	}
}