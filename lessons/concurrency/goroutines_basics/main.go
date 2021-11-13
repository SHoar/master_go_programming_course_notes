package main

import (
	f "fmt"
	"runtime"
	"time"
)

func f1() {
	f.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		f.Println("f1, i=", i)
	}
	f.Println("f1 (goroutine) execution finished")

}

func f2() {
	f.Println("f2 (goroutine) execution started")
	for i := 5; i > 0; i-- {
		f.Println("f2, i=", i)
	}
	f.Println("f2 (goroutine) execution finished")

}

func main() {
	f.Println("main execution started")
	f.Println("No. of cps:", runtime.NumCPU())
	f.Println("No. of Goroutines:", runtime.NumGoroutine())

	f.Println("OS:", runtime.GOOS)
	f.Println("Arch:", runtime.GOARCH)

	f.Println("No. of max Goroutines", runtime.GOMAXPROCS(0))

	go f1()
	f.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	f.Println("main execution stopped")
}
