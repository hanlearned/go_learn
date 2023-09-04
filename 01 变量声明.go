package main

import "fmt"

func main() {
	// 自动推断类型赋值
	var a = 1

	// 声明类型赋值
	var b int = 2

	// 声明变量类型
	var c int
	c = 3

	// 声明赋值
	d := 4

	fmt.Printf("a=%v b=%v c=%v d=%v", a, b, c, d)

}
