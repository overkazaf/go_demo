package main

import "fmt"

func main() {
	a := 4
	b := 4

	c := "abc"
	d := "abc"

	fmt.Println(a == b)

	fmt.Println( &a, &b )

	fmt.Println(c == d)

}