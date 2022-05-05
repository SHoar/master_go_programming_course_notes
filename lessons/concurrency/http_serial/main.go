package main

import (
	f "fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
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

}

func main() {
	urls := []string{"https://golang1.org", "https://medium.com", "https://google.com/a.html"}
	
	for _, url := range urls {
		checkAndSaveBody(url)
		f.Println(strings.Repeat("#", 20))
	}
}
