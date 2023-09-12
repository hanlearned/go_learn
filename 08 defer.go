// 延时函数， 使用 defer 时，被使用函数就会在 return 前结束
package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	for i := range arr1 {
		defer fmt.Println(i)
	}
}
