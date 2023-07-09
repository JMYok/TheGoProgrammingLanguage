package main

import "fmt"

func main() {
	var a [3]int
	var b [3]int = [3]int{1, 2, 3}
	var c = [3]int{4, 5, 6}
	for i, val := range a {
		fmt.Printf("%d:%d\n", i, val)
	}
	for i, val := range b {
		fmt.Printf("%d:%d\n", i, val)
	}
	for i, val := range c {
		fmt.Printf("%d:%d\n", i, val)
	}
}
