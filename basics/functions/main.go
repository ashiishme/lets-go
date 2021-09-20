package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

// for more check packages/basicMath

func main() {
	fmt.Println(greet("Ashish"))
}