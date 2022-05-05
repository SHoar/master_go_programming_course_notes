package main

import (
	"fmt"
	"strings"
)

// variadic == func may take 0 OR MORE arguments

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(a ...float64) (float64, float64) {
	sum  := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	
	returnStr := fmt.Sprintf("Age %d, Full Name: %s", age, fullName)
	return returnStr
}

func main() {
	f1(1, 2, 3, 5)
	f1() // when called with no passed args, but func reqs passed variadic args, returns []int(nil)

	nums := []int{1,2}
	nums = append(nums, 3,4, 5) // example of a variadic func

	f1(nums...)
	f2(nums...)
	fmt.Println("Nums: ", nums)
	
	s, p := sumAndProduct(2.0, 5., 10., 5.6, 5.6, 87.3)
	fmt.Printf("Sum: %#v\nProduct: %#v\n", s, p)

	famousComposer := personInfo(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(famousComposer)

}
