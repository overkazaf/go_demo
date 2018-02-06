package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func scan(dirname string, level int, limit int) {
	fmt.Println("d", dirname)
	if level > limit {
		return
	}
	dir, err := ioutil.ReadDir(dirname)
	if err != nil { panic(err) }

	var dirArray []string
	var idx int = 0
	for _, fi := range dir {
		if fi.IsDir() {
			nextDirname := strings.Join([]string{dirname, fi.Name()}, "/")
			// scan(nextDirname, level+1, limit)
			dirArray = append(dirArray, nextDirname)
			idx++
		} else {
			sep := strings.Repeat(" ", level)
			fmt.Println("-", sep, fi.Name())
		}
	}

	l := len(dirArray)
	if l > 0 {
		for i := 0; i < l; i++ {
			scan(dirArray[i], level+1, limit)
		}
	}
}

func main() {
	dirname := "."
	scan(dirname, 0, 10)
}