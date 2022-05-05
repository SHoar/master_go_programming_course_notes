package main

import f "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	myFriends := make([]string, 4)
	myFriendCount := copy(myFriends, friends)
	_ = myFriendCount // cuz who counts their friends?

	f.Printf("%g\n", friends)
	f.Printf("%g\n", myFriends)
	myFriends = myFriends[0:3] // Diana is no longer a friend
	friends[0] = "Margaret"

	f.Println(friends)
	f.Println(myFriends)

}
