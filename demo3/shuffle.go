package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(arr []int, a int, b int) {
	fmt.Println("swapping index ==> ", a, b)
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

func shuffle(arr []int, l int) []int {
	tNow := time.Now()  
    //时间转化为string，layout必须为 "2006-01-02 15:04:05"  
	// timeNow := tNow.Format("2006-01-02 15:04:05")
	timestamp := tNow.UTC().UnixNano()
	fmt.Println("timestamp", timestamp)
	rand.Seed(timestamp)

	newArr := arr[:]
	for l > 0 {
		idx := rand.Intn(l)
		l--
		if idx != l {
			swap(newArr[:], idx, l)
		}
	}

	return newArr
}

func shuffle2(arr *int, l int) {
	fmt.Println(*arr)
}

func main() {
	arr := []int {2, 3, 4, 5, 6, 7}
	bArray := shuffle(arr, 6)
	fmt.Println("bArray", bArray, arr, &arr, &bArray)

	// shuffle2(&arr, 6)
}