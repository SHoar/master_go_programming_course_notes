package main

import "fmt"

func main() {
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	// goto todo
	// 	x := 5
	// todo:
	// 	fmt.Println("something here")

	// USE GOTO STATEMENTS AT YOUR OWN RISK. SHOULD NOT BE USED MUCH, CAN BE DANGEROUS!

	// x := 1
	// loop2:
	// 	if x != 5 {
	// 		fmt.Println(x)
	// 		x++
	// 		goto loop2
	// 	}

}
