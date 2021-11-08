package main

import f "fmt"

func main() {
	x := 10.10
	f.Println("x mem addr:", &x) // should be the mem alloc addr
	ptr := &x
	f.Println("ptr:", ptr)
	f.Println("addr of ptr & value of x:", &ptr, *ptr)
}