package main
import (
	"os"
	"log"
)

func main() {
	var filename string = "aaa.txt"
	fd, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
}