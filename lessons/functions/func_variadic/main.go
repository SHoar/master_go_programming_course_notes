package main

import "fmt"

// variadic == func may take 0 OR MORE arguments

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func main() {
	f1(1, 2, 3, 5)
	f1()
}
