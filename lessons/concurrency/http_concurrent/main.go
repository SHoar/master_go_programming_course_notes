package main

import (
	f "fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		f.Println(err)
		f.Printf("%s is DOWN\n", url)
	} else {
		defer resp.Body.Close()
		f.Printf("%s => Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				f.Printf("Err1: %s\n", err)
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"
			f.Printf("Writing resp body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				f.Println("Err2:", err)
			} else {
				f.Printf("%s created\n", file)
			}
		}
	}
	wg.Done()

}

func main() {
	urls := []string{"https://golang.org", "https://medium.com", "https://google.com"}

	var wg sync.WaitGroup

	wg.Add(len(urls))
	
	for _, url := range urls {
		f.Println(strings.Repeat("#", 20))
		go checkAndSaveBody(url, &wg)
	}

	f.Println("No. of Goroutines", runtime.NumGoroutine())
	
	wg.Wait()
}
