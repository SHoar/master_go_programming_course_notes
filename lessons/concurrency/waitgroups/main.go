package main

import (
	f "fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	f.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		f.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	f.Println("f1 (goroutine) execution finished")
	defer wg.Done() // or (*wg).Done()
}

func f2() {
	f.Println("f2 (goroutine) execution started")
	for i := 5; i > 0; i-- {
		f.Println("f2, i=", i)
	}
	f.Println("f2 (goroutine) execution finished")

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go f1(&wg)
	f.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())

	f2()

	wg.Wait()
	// time.Sleep(time.Second * 2)
	f.Println("main execution stopped")
}
