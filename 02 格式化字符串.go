package main

import (
	"fmt"
	"strings"
)

func main() {

	head := "http://%v"
	name := "wwww.baidu.com"

	// 拼接字符串: 将 head 和 name 拼接到一起。
	url := fmt.Sprintf(head, name)
	fmt.Println(url) //

	s1 := "abc"
	var s2 string = "30"
	s3 := "xyz"

	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	build.WriteString(s3)

	str4 := build.String()
	fmt.Println(str4)

}
