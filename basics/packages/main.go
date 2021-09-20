package main

import (
	"fmt"
	"math"

	"github.com/ashiishme/lets-go/basics/packages/basicMath"
	"github.com/ashiishme/lets-go/basics/packages/stringUtil"
)

func main() {
	fmt.Println("Floor: ", math.Floor(2.7))
	fmt.Println("Ceil: ", math.Ceil(2.7))
	fmt.Println("Sqrt: ", math.Sqrt(16))

	fmt.Println("Reverse: ", stringUtil.Reverse("Hello"))

	// custom Math packages to do simple maths

	fmt.Println("Add: ", basicMath.Add(2, 2))
	fmt.Println("Sub: ", basicMath.Subtract(10, 5))
	fmt.Println("Multi: ", basicMath.Multiply(10, 10))
	fmt.Println("Divide: ", basicMath.Divide(100, 10))
	fmt.Println("Mod: ", basicMath.Remainder(11, 2))
}