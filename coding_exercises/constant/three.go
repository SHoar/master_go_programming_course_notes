package main

import "fmt"

func main() {
	const secPerDay int64 = 24 * 60 * 60 // hrs, minutes, seconds
	const daysYear = 365

	fmt.Println(daysYear * secPerDay)
}
