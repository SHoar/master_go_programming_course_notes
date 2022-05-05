package main

import f "fmt"

func main() {
	s1 := "Hello, Go!"
	f.Println(s1)

	f.Println("He says\"Hello!\"")

	f.Println(`He says: "Hello!"`)

	s2 := `I like \n Go!`
	f.Println(s2)

	// multiline strings
	f.Println(`
Price: 100
Brand: Nike
	`)

	f.Println(`C:\Users\Sean`)
	f.Println("C:\\Users\\Sean")

	var s3 = "I love " + "Go " + "Programming"
	f.Println(s3 + "!")

	f.Println("El at index 0:", s3[0])
	f.Printf("El at index 0: %#v\n", s3[0])

	// s3[5] = 'x' // strings are immutable, cannot do

}
