package main

import "fmt"

func main() {
	m1 := new(map[int]int)
	m2 := make(map[int]int)

	_, x := (*m1)[0]
	fmt.Println(x)

	if val, ok := m2[0]; ok {
		fmt.Println(val)
	}
}
