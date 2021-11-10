package main

import f "fmt"

type car struct {
	brand string
	price int
}

// method declarations are not permitted on named types that are themselves pointer types
type distance *int

// func (d *distance) m1() {
// 	f.Println("just a message")
// }
// ^^

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

// it changes the value it works on
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
	// the changes are not seen to the outside world (pass by value)
}

// declaring a method with a pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	f.Println(myCar)
	myCar.changeCar1("Renault", 21000)
	f.Println(myCar)
	myCar.changeCar2("Renault", 21000)
	f.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	f.Println(*yourCar)

	yourCar.changeCar2("VW", 30000)
	f.Println(*yourCar)

	(*yourCar).changeCar2("BMW", 45000)
	f.Println(*yourCar)

	f.Println(myCar)
}
