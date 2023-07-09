package main

import (
	"fmt"
	"strings"
	"time"
)

func test1() {
	start := time.Now()
	var s, sep string
	for i := 1; i < 1e6; i++ {
		s += sep + "test"
		sep = " "
	}
	fmt.Printf("%dms elapsed\n", time.Since(start).Milliseconds())
}

func test2() {
	start := time.Now()
	// 字符串数组
	var s [1e4]string
	for i := 1; i < 1e6; i++ {
		s[i] = "test"
	}
	strings.Join(s[1:], " ")
	fmt.Printf("%dms elapsed\n", time.Since(start).Milliseconds())
}

func main() {
	test1()
	test2()
}
