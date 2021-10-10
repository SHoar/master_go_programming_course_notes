package main

import f "fmt" // permitted

const done = false //package scoped
var b int = 10

func main() {
	var task = "Running" // block/local scoped
	f.Println(task, done)

	const done = true // // local scoped
	f.Println("b local", b)
	f.Printf("done in main() is %v\n", done)
	f1()

	f.Println("Bye bye")
}

func f1() {
	f.Printf("done in f1(): %v\n", done) // package scope
	f.Println("b in package", b)
	a := 10
	_ = a
}
