package main

import "fmt"

type Person interface {
	say() string
}

type SinglePerson struct {
}

func (person SinglePerson) say() string {
	return "I am a single person"
}

func main() {
	var p Person
	p = new(SinglePerson)

	var s string = p.say()
	fmt.Println(s)
}