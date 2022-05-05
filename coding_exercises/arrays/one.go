package main

import f "fmt"

func main() {
	// 1
	var cities [2]string
	f.Printf("%#v\n", cities)

	// 2
	grades := [3]float64{85.0, 97.0}
	f.Println(grades)

	// 3
	taskDone := [...]bool{}
	// taskDone[0] = true // cannot assign anything to taskDone because empty array with ellipsis put length at 0.
	// taskDone[1] = false
	f.Printf("taskDone: %#v\n", taskDone)

	// 4
	for i := 0; i < len(cities); i++ {
		f.Printf("index: %v, city: %v\n", i, cities[i])
	}

	// 5
	for i := range grades {
		f.Printf("index: %v, grade: %v\n", i, grades[i])
	}
}
