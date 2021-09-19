package main

import "fmt"

func main() {

	var name string = "Ashish Yadav"
	var number int = 123
	var isBoolean bool = true

	shortHandString := "Software Engineer"

	fmt.Println("String: ", name)
	fmt.Println("Number: ", number)
	fmt.Println("Boolean", isBoolean)
	fmt.Println("Short Hand Declration: ", shortHandString)
	// print type of the variable
	fmt.Printf("%T\n", shortHandString)
}
