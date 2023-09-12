package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}
var w = sync.WaitGroup{}
var numSafe = 0
var numUnSafe = 200000

func AddNum() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		numSafe++
	}
	lock.Unlock()
	w.Done()
}

func SubNum() {
	for i := 0; i < 100000; i++ {
		numUnSafe--
	}
	w.Done()
}

func main() {
	w.Add(2)
	go AddNum()
	go AddNum()

	w.Add(2)
	go SubNum()
	go SubNum()

	w.Wait()
	fmt.Println(numSafe)
	fmt.Println(numUnSafe)

}
