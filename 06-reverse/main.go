package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {

	var str string

	fmt.Println("String to Reverse")
	fmt.Print("Value -> ")
	fmt.Scan(&str)

	revstr := reverse(str)

	palindrome := str == revstr
	fmt.Println(revstr)
	fmt.Println("It is palidrome : ", palindrome)

	// s.Push("hello")
	// s.Push("world")

	// i := 0
	// for i <= s.Len() {

	// 	fmt.Println(s.Pop())
	// 	i += 1
	// }
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func reverseChars(str string) (result string) {
	var revstr string
	strToChars := []rune(str)
	simstack := new(stack.Stack)
	for i := 0; i < len(strToChars); i++ {
		simstack.Push(strToChars[i])
	}

	for i := 0; i < len(strToChars); i++ {
		// fmt.Println(string(strToChars[i]))
		val := string(simstack.Pop().(int32))
		revstr = revstr + val
	}

	return revstr
}
