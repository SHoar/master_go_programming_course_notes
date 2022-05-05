package main

import (
	f "fmt"
	"time"
)

func main() {
	c1 := make(chan int) // unbuffered channel

	// c2 := make(chan int, 3) // buffered channel

	go func(c chan int) {
		f.Println("func goroutine starts sending data to channel")
		c <- 10
		f.Println("after sending data into channel")
	}(c1)

	f.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	f.Println("main goroutine starts receiving data")
	d := <- c1
	f.Println("main goroutine received data:", d)

	time.Sleep(time.Second)
}
