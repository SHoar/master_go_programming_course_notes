package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed int

	var s1 speed = 10
	var s2 speed = 20

	truth := s1 > s2

	var x uint
	x = uint(s1)
	var s3 speed = speed(x)
	_ = s3

	fmt.Println("s1 is greater than s2", truth)

	var parisToLondon km = 465

	distanceInMiles := mile(parisToLondon)
	fmt.Println(distanceInMiles)
}
