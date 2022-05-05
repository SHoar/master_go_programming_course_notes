package main

import (
	f "fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	f.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("lessons/files_basics/a.txt")
	if err != nil {
		// f.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}

	err = os.Truncate("lessons/files_basics/a.txt", 0) // truncate === wipe everything in file but leave file standing

	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	// file, err := os.Open("lessons/files_basics/a.txt") // readOnly
	file, err := os.OpenFile("lessons/files_basics/a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("lessons/files_basics/a.txt")
	p := f.Println

	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is Directory?:", fileInfo.IsDir())
	p("Permissions", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%v does not exist!\n", "b.txt")
		}
	}

	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
