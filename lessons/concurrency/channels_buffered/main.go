package main

import (
	f "fmt"
	"time"
)

func main() {
	// c1 := make(chan int) // unbuffered channel

	c2 := make(chan int, 3) // buffered channel

	go func(c chan int) {
		for i:= 1; i <= 5; i++ {
			f.Printf("func goroutine #%d starts sending data to channel\n", i)
			c2 <- i
			f.Printf("func goroutine #%d after sending data into channel\n", i)
		}
		close(c2)
	}(c2)

	f.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c2 { // v := <- c2
		f.Println("main goroutine received data from channel", v)
	}
	
	f.Println("main goroutine sleeps again for 1 second", <- c2)
	time.Sleep(time.Second)
}
