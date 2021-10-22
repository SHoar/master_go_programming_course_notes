package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	mascot := "The Go gopher is an iconic mascot!"
	written, err := file.Write([]byte(mascot))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(written)

}