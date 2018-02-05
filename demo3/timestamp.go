package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	t := time.Now().UTC().UnixNano()
	str := strconv.FormatInt(t, 16)
	fmt.Println(t)

	fmt.Println(str)
}