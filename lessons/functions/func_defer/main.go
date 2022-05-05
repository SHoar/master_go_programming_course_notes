package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("this is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	foo()
	bar()
}

func barfoo() {
	defer foo()
	bar()

	fmt.Println("Just a string after deferring foo() and calling bar()")
}

func main() {
	barfoo()
	defer foobar()

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Opened file main.go")
	defer file.Close()
}
