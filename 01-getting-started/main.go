package main

import (
	"fmt"
)

func main() {

	// Primitive data types
	var i int = 12
	var j float32 = 12.5
	var h bool = true
	var str string = "Welcome"
	var com complex64 = 1 + 2i

	// Printf in golang
	fmt.Printf("Type %T Value : %v  \n", i, i)
	fmt.Printf("Type %T Value : %v  \n", j, j)

	// Println to console
	fmt.Println("Integer Value = ", i)
	fmt.Println("Float Value = ", j)
	fmt.Println("Boolean Value = ", h)
	fmt.Println("String Value = ", str)
	fmt.Println("Complex Value = ", com)

	// multiple assignment
	re, im := real(com), imag(com)

	
	fmt.Println("Complex Real Number = ", re)
	fmt.Println("Complex Imaginary Number = ", im)

}
