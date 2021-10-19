package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// readfile with os package
	// opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	// write file with file package
	bytesWritten, err := file.Write(byteSlice) // os.Write returns the no. of bytes written and an error value
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", bytesWritten)

	bs := []byte("Go Programming is cool!")
	// write file with ioutil and set permissions
	err = ioutil.WriteFile("c.txt", bs, 0644)
	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.
	if err != nil {
		log.Fatal(err)
	}
}
