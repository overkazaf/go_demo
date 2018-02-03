package main

import "fmt"

func main() {
	var s string
	var flag bool = true
	for flag {
		fmt.Scanf("%s", &s)

		fmt.Println("You input ", s)

		if s == "done" {
			flag = false
		} else {
			flag = true
		}
	}
}
