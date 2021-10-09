package main

import f "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}

	myFriends := append(friends, "Jim")
	friends[0] = "Mary"

	f.Println(friends)
	f.Println(myFriends)
}
