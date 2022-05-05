package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main()  {

	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Could not read file")
	}

	log.Printf("Data as string: %s\n", data)
}