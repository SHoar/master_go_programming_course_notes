package main

import f "fmt"

func swapFLoatValues(x *float64, y *float64)  {
	*x, *y = *y, *x	
}

func main() {
	x, y := 5.5, 8.8
	f.Println("x, y BEFORE swap:", x, y)
	swapFLoatValues(&x, &y)
	f.Println("x, y AFTER swap:", x, y)
}