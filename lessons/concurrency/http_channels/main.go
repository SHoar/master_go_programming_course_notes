package main

import (
	f "fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := f.Sprintf("%s is DOWN\n", url)
		s += f.Sprintf("Error: %s\n", err)
		f.Println(s)
	} else {
		s := f.Sprintf("%s => Status Code: %d \n", url, resp.StatusCode)
		s += f.Sprintf("%s is UP\n", url)
		f.Println(s)
	}
	c <- url // sending into the channel

}

func main() {
	urls := []string{"https://golang.org", "https://medium.com", "https://google.com"}

	c := make(chan string)

	// for _, url := range urls {
	// 	go checkUrl(url, c)
	// }

	f.Println("No. of Goroutines", runtime.NumGoroutine())

	// for i:=0; i< len(urls); i++ {
	// 	f.Println(strings.Repeat("#", 20))
	// 	f.Println(<-c)
	// }

	// 1. via infinite loop - continue to check url from channel every 2 secs
	// for {
	// 	go checkUrl(<- c, c)
	// 	f.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// 2. via controlled for loop
	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkUrl(url, c)
	// }

	// 3. via anonymous func
	for url:= range c {
		go func(u string){
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
	}

}
