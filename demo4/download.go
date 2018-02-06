package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"io"
)

func check_and_log(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var url string = "https://ss1.bdstatic.com/5eN1bjq8AAUYm2zgoY3K/r/www/cache/static/protocol/https/home/img/qrcode/zbios_efde696.png"
	response, e := http.Get(url)
	check_and_log(e)
	defer response.Body.Close()

	var filename string = "download.png"
	fd, err := os.Create(filename)
	check_and_log(err)

	b, err := io.Copy(fd, response.Body)
	check_and_log(err)

	fd.Close()

	fmt.Println("download png file successfully..., file size:", b)
}