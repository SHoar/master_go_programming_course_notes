package main

import f "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbors := friends

	friends["Dan"] = 50

	f.Println(neighbors) // neighbors has pointers to friends map

	people := make(map[string]int)

	// how to clone a map in GO
	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18 // not updated in people b/c people is a separate clone
	f.Printf("P %#v ------------------ F %#v\n", people, friends)

	friends["Maria"] = 29

	f.Printf("%#v ------------------%#v\n", people, friends)

	delete(friends, "Dan")

	f.Printf("%#v ------------------%#v\n", people, friends)
	delete(people, "ANDREI") // no error even if key doesn't exist !!!
}
