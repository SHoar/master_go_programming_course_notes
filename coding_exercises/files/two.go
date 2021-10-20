package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err){
			log.Println("File does not exist")
		}
		log.Fatal(err)
	}
	
	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal(err)
	}

}
	