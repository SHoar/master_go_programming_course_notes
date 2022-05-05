package main

import (
	f "fmt"
	"strings"
)

func main() {
	p := f.Println
	result := strings.Contains("I love Go Programming!", "lovex")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO pYTHON jAVA"))
	p(strings.ToUpper("GO pYTHON jAVA"))

	p("go" == "go")

	p("GO" == "go")

	// inefficient way
	p(strings.ToLower("GO") == strings.ToLower("go"))

	p(strings.EqualFold("Go", "go")) // compares strings regardless of upper/lower case

	myStr := strings.Repeat("ab", 10)
	p(myStr)
	myStr = strings.Replace("192.168.0.1", ".", ":", 2)
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	f.Printf("%T\n", s)
	f.Printf("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	f.Printf("%#v\n", s)
	f.Println(strings.Join(s, ""))

	s = []string{"I", "am", "learning", "Golang"}
	myStr = strings.Join(s, " ")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	colorFields := strings.Fields(myStr)
	p(colorFields)

	f.Printf("%T\n", colorFields)
	f.Printf("%#v\n", colorFields)

	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!")
	f.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello Gophers!!!?", ".!?")
	f.Printf("%q\n", s2)
}
