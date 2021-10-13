package main

import f "fmt"

func main() {
	var employees map[string]string
	f.Printf("%#v\n", employees)
	f.Printf("No. of pairs: %d\n", len(employees))

	f.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	// employees["name"] = "Dan"
	// employees["role"] = "Supervisor"
	// employees["eID"] = "1001"

	var accounts map[string]float64
	f.Printf("%#v\n", accounts["0x323"])

	// slices cannot be keys, but arrays can
	var m1 map[[5]int]string
	_ = m1

	people := map[string]float64{}
	people["John"] = 20.4
	people["Mary"] = 10

	f.Println(people)

	for k, v := range people {
		f.Printf("The value for key %q is %v\n", k, v)
	}

	map1 := make(map[int]int)
	map1[4] = 8

	f.Println(map1)

	balances := map[string]float64{
		"USD": 323.11,
		"GBP": 400.54,
		"EUR": 432.1,
		"RON": 500.01,
	}

	f.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	f.Println(m)

	v, ok := balances["RON"]
	if ok {
		f.Println("RON = ", v)
	} else {
		f.Println("The Ron key doesn't exist in the map")
	}

	for k, v := range balances {
		f.Printf("Key: %#v, Value: %#v\n", k, v)
	}
}
