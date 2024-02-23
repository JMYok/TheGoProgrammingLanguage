package main

import (
	"fmt"
	"sync"
)

func main() {
	m := new(sync.Map)
	m.Store("今天", "对，今天")
	m.Store("明天", "明日复明日")
	value, ok := m.Load("今天")
	if ok {
		fmt.Println(value)
	}
	f := func(key, value interface{}) bool {
		fmt.Printf("遍历：key:%v,value:%v\n", key, value)
		return true
	}
	m.Range(f)
}
