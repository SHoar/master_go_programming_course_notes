package main

import f "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	myFriends := make([]string, 4)
	myFriendCount := copy(myFriends, friends)
	_ = myFriendCount // cuz who counts their friends?

	myFriends = myFriends[0:3] // Diana is no longer a friend

	f.Println(friends)
	f.Println(myFriends)

}
