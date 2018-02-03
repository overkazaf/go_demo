package main

import "fmt"

func main() {
	// array
	var a = [...]int {0, 1, 2, 3, 4}
	
	var b = [...][4]int {
		{ 1, 2, 3, 4 },
		{ 3, 4, 5, 6 },
		{ 6, 7, 8, 9 },
	}
	// loop
	for i := 0; i < 4; i++ {
		fmt.Println(a[i])
	}

	for x := 0; x < 3; x++ {
		for y := 0; y < 4; y++ {
			fmt.Println(b[x][y])
		}
		println()
	}
}