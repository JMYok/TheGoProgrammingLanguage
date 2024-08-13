package main

import (
	"fmt"
	"unsafe"
)

func main() {
	ch := make(chan int, 1)
	fmt.Printf("Channel address: %p\n", unsafe.Pointer(&ch))
	//ch <- 1
	//fmt.Println(<-ch)
}
