package main

import "fmt"

func main() {

	// string
	var name string = "Ashish Yadav"
	// integer
	var number int = 123
	// boolean
	var isBoolean bool = true

	// shorthand
	shortHandString := "Software Engineer"

	// multiple shorthand assginment [ This is interesting ]
	contactNumber, address := 9877712355, "Butwal"

	fmt.Println("String: ", name)
	fmt.Println("Number: ", number)
	fmt.Println("Boolean", isBoolean)
	fmt.Println("Short Hand Declration: ", shortHandString)
	// print type of the variable
	fmt.Printf("%T\n", shortHandString)

	fmt.Println("Contact: ", contactNumber)
	fmt.Println("Address: ", address)
}
