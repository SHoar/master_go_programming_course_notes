package main

import "fmt"

func main(){
	type Contact struct {
		email string
		address string
		phone  int
	}

	type Employee struct {
		name string
		salary int
		contactInfo Contact
	}

	john := Employee{
		name: "John Kellogg", 
		salary: 40000, 
		contactInfo: Contact{
			email: "jkeller@company.com",
			phone: 555123456789, 
			address: "20 Street, London, UK",
		},
}


	fmt.Printf("%v\n", john)

	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)
	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

	myContactInfo := Contact{
		email: "curly_brace@somewhere.com",
		address: "1234 Somewhere street, North Something, MD",
		phone: 5551234567890,
	}

	fmt.Println(myContactInfo)

	john.contactInfo = myContactInfo

	fmt.Println(john)

}