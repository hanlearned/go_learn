// 值类型和引用类型
package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

// * 声明指针变量
// & 指针变量的存储地址
// * 取值运算符
// & 取地址运算符
func swap(a *int, b *int) {
	// a *int 是声明一个形参为 a 的int类型的指针变量， 传入时需要传入地址。
	tmp := *a // 此时取出a的值赋值给tmp
	*a = *b
	*b = tmp
	return
}

func main() {
	first := 10
	second := 20
	swap(&first, &second)
	fmt.Println("first=", first)   // first= 20
	fmt.Println("second=", second) // second= 10
}
