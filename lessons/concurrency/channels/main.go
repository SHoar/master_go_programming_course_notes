package main

import f "fmt"

func main() {
	var ch chan int
	f.Println(ch)

	ch = make(chan int)
	f.Println(ch)

	c := make(chan int)
	// <<- channel operator

	//SEND
	c <- 10

	// RECEIVE
	num := <- c

	f.Println(<- c)


	_ = num

	close(c)
}