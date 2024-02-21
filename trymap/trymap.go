package main

import "fmt"

func main() {
	var m1 map[int]struct{}
	var m2 = make(map[int]struct{}, 0)
	s := m1[1]
	fmt.Println(s)
	//nil map进行写
	//panic: assignment to entry in nil map
	m1[2] = struct{}{}

	s2 := m2[1]
	fmt.Println(s2)
}
