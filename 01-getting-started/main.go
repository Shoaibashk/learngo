package main

import (
	"fmt"
)

func main() {

	var i int = 12
	var j float32 = 12.5
	var h bool = true
	var str string = "Welcome"
	var com complex64 = 1 + 2i

	fmt.Println("Integer Value = ", i)
	fmt.Println("Float Value = ", j)
	fmt.Println("Boolean Value = ", h)
	fmt.Println("String Value = ", str)
	fmt.Println("Complex Value = ", com)
	fmt.Println("Complex Real Number = ", real(com))
	fmt.Println("Complex Imaginary Number = ", imag(com))

}
