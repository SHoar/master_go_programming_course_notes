package main

import "fmt"

func main() {
	type book struct {
		title string
		author string
		year int
	}

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	// fmt.Println(lastBook.published) // what if we ask for something that's not there

	fmt.Printf("%v\n", lastBook)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878

	fmt.Printf("%+v\n", lastBook)

	// comparing structs
	aBook := book{title: "Anna Karenina", author: ""}

	fmt.Println(aBook.title == lastBook.title)

	newBook := lastBook
	newBook.year = 2020

	fmt.Println(newBook, lastBook)
}