package main

import "fmt"

func main() {
	a := nextId()
	fmt.Println(a());
	fmt.Println(a());
	fmt.Println(a());
	fmt.Println(a());
}

func nextId() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}