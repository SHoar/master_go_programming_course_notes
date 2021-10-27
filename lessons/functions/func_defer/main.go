package main

import "fmt"

func foo(){
	fmt.Println("this is foo()")
}

func bar(){
	fmt.Println("This is bar()")
}

func foobar(){
	foo()
	bar()
}

func barfoo(){
	defer foo()
	bar()
}

func main()  {
	barfoo()
}