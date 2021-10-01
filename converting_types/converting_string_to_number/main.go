package main

import (
	"fmt"
	"strconv"
)

func main() {
	// number to string
	s := string(99) //ascii code for c
	fmt.Println(s)

	s = fmt.Sprint(99)
	fmt.Println(s)

	// s1 := string(44.2) //connot convert untyped float constant to string
	myStr := fmt.Sprintf("%.1f", 44.2) //set decimal places, and use fmt to print num to string
	fmt.Println(myStr)

	myStr1 := fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	// ascii/unicode
	fmt.Println(string(34234))

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	// import strconv package for string conversions
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err // catch pototential errs
	fmt.Println(f1)

	i, err := strconv.Atoi("-50") // Ascii to int
	_ = err
	s2 := strconv.Itoa(20)        // Int to Ascii

	fmt.Printf("i type is %T, value is %v\n", i, i)
	fmt.Printf("s2 type is %T, value is %q", s2, s2)
}
