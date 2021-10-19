package main

import f "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}

	sa := f.Sprintf("%s", a)
	sb := f.Sprintf("%s", b)

	f.Println(sa, sb)

	if sa == sb {
		f.Println("Maps are equal")
	} else {
		f.Println("Maps are not equal")
	}
}
