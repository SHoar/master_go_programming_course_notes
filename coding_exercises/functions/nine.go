package main

import "fmt"

func printInt(n int) {
	arg := n
	fmt.Printf("Variable: %v, Type: %T\n", arg, arg)
}

func main() {
	printed := printInt
	printed(5)
	printed(9)
}