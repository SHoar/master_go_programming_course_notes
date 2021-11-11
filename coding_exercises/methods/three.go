package main

import f "fmt"

type book struct {
	price float64
	title string
}

func (b *book) discount() {
	(*b).price = (*b).price - ((*b).price * 0.10)
}

func main() {
	hp := book{price: 35.99, title: "Harry Potter"}
	f.Printf("Orig. price: %.2f\n", hp.price)
	hp.discount()
	f.Printf("Discount price: %.2f\n", hp.price)
}