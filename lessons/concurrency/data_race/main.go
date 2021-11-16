package main

import (
	f "fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0

	for i:= 0; i<gr; i++ {
		go func ()  {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	f.Println("Final value of n:", n)
}

// test data races with `go run -race main.go`