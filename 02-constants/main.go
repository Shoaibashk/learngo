package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHeadQuarter // 1^1 * 1^1 = 1 ^ 2
	isManager // 1^1 * 1^2 = 
	isDevOps
	isItCrowd
	isHuman
)

func main() {
	var roles byte = isAdmin | isManager | isHeadQuarter
	fmt.Println(roles)
	fmt.Println(roles & isAdmin)
	fmt.Println(roles&isManager == isManager)

}
