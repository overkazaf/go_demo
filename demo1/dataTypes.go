package main

import "fmt"

/**
	1. 布尔型
		bool
	2. 数字型
		byte
		uint8, uint16, ..., uint32
		int8, int16, ..., int32
		float32, float64
		complex64, complex128
	3. 字符串
	4. 派生类型
		(a) 指针类型（Pointer）
		(b) 数组类型
		(c) 结构化类型(struct)
		(d) Channel 类型
		(e) 函数类型
		(f) 切片类型
		(g) 接口类型（interface）
		(h) Map 类型
*/

func main() {
	var a int
	var b float32

	a = 3
	b = 2.3

	fmt.Println("%d, %2f", a, b)
}