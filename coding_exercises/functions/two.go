package main

import f "fmt"

func f1(n uint) (int, int) {
	factorial := 1
	sum := 0
	for i := 1; i <= int(n); i++ {
		factorial *= i
		f.Printf("%v\n", i)
		// if i > 0 && i <= n {
		sum += i
		// }
	}

	return factorial, sum
}

func main() {
	factorial, sum := f1(5)
	f.Printf("F: %v, S: %v\n", factorial, sum)
}
