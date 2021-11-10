package main

import (
	f "fmt"
	"log"
	"strconv"
	"strings"
)

func myFunc(n string) int {
	var sum int
	converted := []string{}
	for _, dig := range [3]int{} {
		dig, _ = strconv.Atoi(n)
		if dig == 1 {
			converted[dig] = string(dig)
			sum += dig
		}
		if dig > 1 {
			digStr := strings.Repeat(n, dig)
			converted[dig] = digStr
			adder, err := strconv.Atoi(converted[dig])
			if err != nil {
				log.Fatal(err)
			}
			sum += adder
		}

	}

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = dig
	// dig2str := strings.Repeat(n, 2)
	// dig3str := strings.Repeat(n, 3)
	// dig2, err := strconv.Atoi(dig2str)
	// dig3, err := strconv.Atoi(dig3str)

	// converted := []int{dig, dig2, dig3}
	// for _, i := range converted {
	// 	sum += i
	// }

	return sum
}

func main() {
	f.Println("myFunc(5):", myFunc("5"))
}
