package main

import (
	"fmt"
	"time"
)

func add(ch chan bool, num *int) {
	ch <- true
	*num = *num + 1
	<-ch
}

func main() {
	ch := make(chan bool, 1)

	var num int
	for i := 0; i < 1000; i++ {
		go add(ch, &num)
	}
	time.Sleep(2)
	fmt.Println("num的值:", num)
}
