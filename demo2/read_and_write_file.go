package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	read()
	write()
	read_and_write()
}

func read() {
	var filename string = "./README.md"
	var f *os.File
	var err error
	var n int

	f, err = os.OpenFile(filename, os.O_RDONLY, 0666)

	if err != nil {
		panic(err)
	}

	fmt.Println("文件存在")

	
	for {
		buffer := make([]byte, 1024)
		n, _ = f.Read(buffer)

		if 0 == n {
			break
		}

		fmt.Println("一共", n, "个字符")
		os.Stdout.Write(buffer[:n])
	}

	f.Close()
}

func write() {
	var filename string = "./write_demo.txt"
	var testing_string string = "testing4444\r\n"

	var fin, fos *os.File
	var _1, _2 error

	fin, _ = os.OpenFile(filename, os.O_WRONLY, 0666)
	defer fin.Close()
	if _1 != nil {
		fmt.Println("文件已经存在")
	} else {
		fmt.Println("文件未存在")
	}

	fos, _2 = os.Create(filename)

	if _2 != nil {
		panic(_2)
	}

	// var buf []byte = []byte(testing_string)
	// var len int = len(testing_string)
	fos.Write([]byte(testing_string))
	defer fos.Close()
}

func check(filename string) bool {
	f, err := os.Open(filename)
	defer f.Close()
	return err == nil
}

func create(filename string) bool {
	f, err := os.Create(filename)
	defer f.Close()

	return err == nil
}

func check_and_create(filename string) bool {
	if check(filename) == false {
		if create(filename) {
			return true
		} else {
			return false
		}
	}
	return true
}

func prepare_files(src string, target string) {
	check_and_create(src)
	check_and_create(target)
}

func read_and_write() {
	var target_file_name string = "target.txt"
	var source_file_name string ="source.txt"

	prepare_files(source_file_name, target_file_name)

	fin, e := os.OpenFile(source_file_name, os.O_RDONLY, 0666)
	if e != nil { panic(e) }

	fout, e := os.OpenFile(target_file_name, os.O_WRONLY, 0666)
	if e != nil { panic(e) }

	for {
		buf := make([]byte, 1024)
		n, err := fin.Read(buf)

		if err != nil && err != io.EOF { panic(err) }
		if 0 == n {
			fmt.Println("err", err)
			fout.WriteString("\nI AM A COPIED FILE\n")
			break
		}

		if n2, err := fout.Write(buf[:n]); err != nil {
			//写入output.txt,直到错误
            fmt.Println("error:::", err)
        } else if n2 != n {
            panic("error in writing")
        }

	}

	defer fin.Close()
	defer fout.Close()

	fmt.Println("File control has been done...")
}

