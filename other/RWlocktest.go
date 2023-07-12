package main

import (
	"fmt"
	"sync"
	"time"
)

var cnt = 0

func main() {
	var mr sync.RWMutex
	for i := 1; i <= 3; i++ {
		go write(&mr, i)
	}
	for i := 1; i <= 3; i++ {
		go read(&mr, i)
	}
	time.Sleep(time.Second)
	fmt.Println("final count:", cnt)
}

func read(mr *sync.RWMutex, i int) {
	fmt.Printf("goroutine%d reader start\n", i)
	mr.RLock()
	fmt.Printf("goroutine%d reading count:%d\n", i, cnt)
	time.Sleep(time.Millisecond)
	mr.RUnlock()
	fmt.Printf("goroutine%d reader over\n", i)
}

func write(mr *sync.RWMutex, i int) {
	fmt.Printf("goroutine%d writer start\n", i)
	mr.Lock()
	cnt++
	fmt.Printf("goroutine%d writing count:%d\n", i, cnt)
	time.Sleep(time.Millisecond)
	mr.Unlock()
	fmt.Printf("goroutine%d writer over\n", i)
}
