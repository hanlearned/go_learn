package main

import "fmt"

func main() {
	// make 的使用
	// make() 是 Go 语言内存分配的内置函数

	map1 := make(map[string]string, 3)
	arr1 := make([]string, 5, 6)
	chan1 := make(chan string, 2)

	// append
	map1["name"] = "韩学成"
	arr1 = append(arr1, "一个字符串")
	chan1 <- "chan1"

	fmt.Printf("map1=%v, arr1=%v, chan1=%v \n", map1, arr1, chan1)
	// map1=map[name:韩学成], arr1=[     一个字符串], chan1=0xa046040

	// delete
	delete(map1, "name")
	fmt.Printf("map1=%v", map1)

}
