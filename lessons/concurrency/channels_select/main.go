package main

import (
	f "fmt"
	"time"
)
func main() {
	start := time.Now().UnixNano() / 1000000
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(2 * time.Second)
		c1 <- "Hello!"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		c2 <- "Salut!"
	}()
	f.Println("Start time:", start)
	for i:=0; i<2; i++{
		select {
		case msg1 := <- c1:
			f.Println("Received:", msg1)
		case msg2 := <- c2:
			f.Println("Received:", msg2)
		}
	}
	end := time.Now().UnixNano() / 1000000
	f.Println("end time", end)
	total := start - end
	
	f.Println("total time:", start - end)
}