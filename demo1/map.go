package main

import "fmt"

func main() {
	var statusMap = map[string] int {
		"ORDER_READY": 1,
		"ORDER_PAYED": 2,
		"ORDER_CANCEL": 3,
		"ORDER_FAILED": 4,
		"ORDER_SUCCESS": 5,
	}

	for ORDER_STATUE := range statusMap {
		fmt.Println("Capital of", ORDER_STATUE,"is", statusMap[ORDER_STATUE])
	 }
}