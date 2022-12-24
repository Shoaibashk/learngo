package main

import "fmt"

func main() {
	grades := [...]int{45, 65, 85}
	fmt.Println(grades)

	basicArray()
	copyArray()
}

func copyArray() {
	a := [3]int{12, 32, 52} // (its a array not a  slice )
	b := a                  // shadow copy to b variable
	// b := &a                 // reference type to b variable

	b[1] = 34
	fmt.Println("A : ", a)
	fmt.Println("B : ", b)
}

func basicArray() {
	var students [3]int

	students[0] = 12
	students[1] = 32
	students[2] = 43

	fmt.Println(students)
}
