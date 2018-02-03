package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	var c float32 = 2.0
	var d float32 = 2.0

	fmt.Println(&a, &b, &c, &d, *(&d))
}