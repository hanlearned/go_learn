package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%v", i)
	}
	fmt.Println()
	wg.Done() // 任务结束登记
}

func main() {
	// 登记 goroutine +1
	wg.Add(1)
	go say()
	wg.Add(2) // 登记 goroutine +2
	go say()
	go say()
	wg.Wait() // 等待所有登记的协程都结束
}
