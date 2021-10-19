package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bytesSlice := make([]byte, 5)
	numberBytesRead, err := io.ReadFull(file, bytesSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", bytesSlice)

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file) // this will read all of the file without having to declare how many bytes to read
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data as a string: %s\n", data)
	log.Printf("Number of bytes read: %d\n", len(data))

	data, err = ioutil.ReadFile("test.txt") // opens, and reads file  all at once
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
	log.Println("Number of bytes read: ", len(data))
}
