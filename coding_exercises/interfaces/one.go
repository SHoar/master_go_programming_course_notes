package main

import f "fmt"

type vehicle interface {
    License() string
    Name() string
}
 
type car struct {
    licenceNo string
    brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	var v vehicle = car{licenceNo: "A1A-8AMA", brand: "Ford"}
	f.Printf("%v\n", v)
	f.Println("Lola's License No:", v.License())
	f.Println("Lola's Make:", v.Name())

}