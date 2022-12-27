package main

import "fmt"

func main() {
	var firstName *string = new(string)

	*firstName = "hello"
	fmt.Println(&firstName)
	fmt.Println(*firstName)
}
