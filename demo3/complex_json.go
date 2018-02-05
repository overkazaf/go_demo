package main

import (
	"fmt"
	// "encoding/json"
)

type Person struct {
	Id      int
	Name    string
	Age     int
	FriendIds []int
}

func (p *Person)showPerson() {
	fmt.Println(p.Id, p.Name, p.Age, p.FriendIds)
}

func run() {
	p1 := &Person {
		1,
		"Kevin",
		12,
		[]int{2, 3, 4},
	}
	p2 := &Person{2, "Tom", 33, []int{1, 3}}
	p3 := &Person{3, "Jim", 23, []int{1, 2}}
	p4 := &Person{4, "Bob", 31, []int{1}}

	p1.showPerson()
	p2.showPerson()
	p3.showPerson()
	p4.showPerson()
}

func main() {
	run()
}