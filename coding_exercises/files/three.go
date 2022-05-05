package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Stat("data.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	os.Remove(file.Name())
}
