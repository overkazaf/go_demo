package main

import "fmt"

func main() {
	a := 3
	b := 5
	fmt.Println(max(a, b))

	c := "hello"
	d := "world"
	fmt.Println(swap(c, d))
}

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func swap(a string, b string) (string, string) {
	return b, a
} 