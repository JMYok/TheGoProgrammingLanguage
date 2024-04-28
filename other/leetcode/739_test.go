package main

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	var st []int
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		top := len(st) - 1
		for len(st) > 0 && t >= temperatures[st[top]] {
			st = st[:top]
		}
		if len(st) <= 0 {
			ans[i] = 0
		} else {
			ans[i] = st[top]
		}
		st = append(st, i)
	}
	return ans
}
