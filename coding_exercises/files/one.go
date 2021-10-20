package main

import (
	"log"
	"os"
)

func main()  {
	fileInfo, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fileInfo.Close()
}