package main

import (
	f "fmt"
	"math"
)

func cubed(a float64) float64 {
	return math.Pow(a, 3)
}

func main() {
	cube := cubed(5)
	f.Println(cube)
}
