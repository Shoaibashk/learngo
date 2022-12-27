package main

import "fmt"

const (
	isAdmin       = 1 << iota // 00000001
	isHeadQuarter             // 00000010
	isManager                 // 00000100
	isDevOps                  // 00001000
	isItCrowd
	isHuman
)

func main() {
	var roles byte = isAdmin | isManager | isHeadQuarter
	fmt.Println(roles)
	fmt.Println(roles & isAdmin)
	fmt.Println(roles&isManager == isManager)
	constantTest()

}

func constantTest() {
	const c = 3
	v := c + 1.2
	fmt.Println(v)
}
