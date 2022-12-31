package main

import "fmt"

type Person struct {
	id   int
	name string
	car  string
}

// func (p Person) Default(id int, name string, car string) {
// 	p.id = id
// 	p.name = name
// 	p.car = car
// }

func main() {
	// var mp map[string]int = map[string]int

	mp := map[string]int{
		"shoaib": 0,
		"ahmed":  1,
	}
	mp1 := make(map[int]Person)

	mp1[0] = Person{id: 0, name: "shoaib", car: "bmw"}
	mp1[1] = Person{id: 1, name: "ahmed", car: "mercedes"}
	mp1[2] = Person{id: 2, name: "covery", car: "tesla"}
	mp1[3] = Person{id: 3, name: "nim", car: "hammer"}

	mp1Val, mp1Ok := mp1[4]

	fmt.Println(mp1Val, mp1Ok)

	fmt.Println(mp1)
	val, ok := mp["shoaib"]

	fmt.Println(val, ok)
	mp["hello"] = 2
	fmt.Println(mp)
}
