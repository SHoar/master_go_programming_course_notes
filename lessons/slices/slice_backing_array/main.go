package main

import (
	f "fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600
	f.Println(s1)
	f.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3] // uses arr1[startIndex: upTo stopIndex] to "back" the slice
	arr1[1] = 2
	f.Println(slice1, slice2)

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...) // because is a slice backing a slice, does not have pointer updates

	cars[0] = "Nissan"
	f.Println(cars, newCars) // because is a slice backing a slice, does not have pointer updates

	//remember s1 ? (ln 6)
	newSlice := s1[2:5]
	f.Println(len(newSlice), cap(newSlice)) // newSlice has len and cap from backing array s1

	f.Printf("%p\n", &s1)               // slice header pointer
	f.Printf("%p %p\n", &s1, &newSlice) // different memory addresses

	newSlice[0] = 1000
	f.Println("s1: ", s1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}
	_ = s
	f.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a))
	f.Printf("array's size in bytes: %d \n", unsafe.Sizeof(s)) // slice occupies less memory size

}
