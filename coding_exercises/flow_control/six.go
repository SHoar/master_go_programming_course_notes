package main

import f "fmt"

func main() {
	age := 19
	switch {
	case age < 0 || age > 100:
		f.Println("Invalid Age")
		break
	case age < 18:
		f.Println("You are minor!")
		break
	case age == 18:
		f.Println("Congratulations! You've just become major!")
		break
	default:
		f.Println("You are major!")
		break
	}
}
