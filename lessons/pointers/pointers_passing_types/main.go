package main

import (
	"fmt"
	"strings"
)

func changeValues(quantity int, price float64, name string, sold bool)  {
	quantity = 3
	price = 500.4
	name = "mobile phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool)  {
	*quantity = 3
	*price = 500.4
	*name = "mobile phone"
	*sold = false
}

type Product struct {
	price float64
	productName string
}

func changeProduct(p Product) {
	p.price = 200
	p.productName = "new name"
}

func changeProductByPointer(p *Product) {
	(*p).price = 200.99
	(*p).productName = "something else"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}


func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println(quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println(quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println(quantity, price, name, sold)

	gift := Product {
		price: 75.50,
		productName: "gift",
	}

	fmt.Println("Gift BEFORE changeProduct:", gift)
	changeProduct(gift)
	fmt.Println("Gift AFTER  changeProduct:", gift)
	changeProductByPointer(&gift)
	fmt.Println("Gift AFTER changeProductByPointer", gift)

	// changeSlice & changeMap
	newSlice := []int{1,2,3,4,5}
	newMap := map[string]int{
		"a": 30,
		"b": 40,
		"c": 50,
	}
	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("newSlice BEFORE changeSlice", newSlice)
	changeSlice(newSlice)
	fmt.Println("newSlice AFTER changeSlice", newSlice)
	
	fmt.Println("newMap BEFORE changeMap", newMap)
	changeMap(newMap)
	fmt.Println("newMap AFTER changeMap", newMap)

}