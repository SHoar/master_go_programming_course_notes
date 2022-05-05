package main

import "fmt"

func main() {
    title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
    // title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

    fmt.Println("Book1", title1, author1, year1)
    fmt.Println("etc...")

    type book struct {
        title string
        author string
        year int
    }

    type book1 struct {
        title, author string
        year int
    }

    mybook := book{title1, author1, year1}

    // mybook1 := book1{title1, author1, year1}

    fmt.Println(mybook)

    bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
    _ = bestBook
}