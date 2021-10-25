package main

import f "fmt"

func main() {
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	f.Printf("%#v\n", diana)
	f.Printf("%v\n", diana)

	f.Printf("Diana's Age: %d\n", diana.age)

	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1984 by George Orwell", 10.2, false}
	f.Printf("%#v\n", b1)

	// just the string part
	f.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"Jon", 400000, false}
	f.Printf("%#v\n", e)

	// mod anonymous bool attribute
	e.bool = true
	f.Printf("%#v\n", e)

}
