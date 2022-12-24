package main

import "fmt"

func main() {
	sliceArray()
}
func sliceArray() {
	// Known as dynamic array
	a := []int{1, 2, 3} // its a slice
	b := a              // its a by default reference types

	fmt.Println(b)
	fmt.Println(a)
}
