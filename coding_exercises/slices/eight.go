package main

import "fmt"

func main() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	yearsLen := len(years)
	var newYears []int
	newYears = append(years[0:3], years[yearsLen-3:]...)

	fmt.Println(years, len(years))
	fmt.Println(newYears, len(newYears))
}
