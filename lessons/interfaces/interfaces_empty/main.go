package main

import f "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}
	empty = 5
	f.Println(empty)

	empty = "Go"
	f.Println(empty)

	empty = []int{4,5,6}
	f.Println(empty)
	f.Println(len(empty.([]int)))

	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5,6,7}
	f.Println(you.info)
}
