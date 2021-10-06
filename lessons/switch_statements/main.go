package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python")
		break
	case "Javascript":
		fmt.Println("use or do not semicolons")
		break
	case "Go", "golang":
		fmt.Println("You will use curly braces")
		break
	default:
		fmt.Println("Any other programming language is being used")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0:
		fmt.Println("Odd")
	default:
		fmt.Println("Never here")
	}

	hour := time.Now().Hour()
	// fmt.Printf("Hour: %v", hour)
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good Afternoon")
	case hour >= 17:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Day!")
	}
}
