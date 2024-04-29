package main

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	var st []int
	n := len(temperatures)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		if len(st) > 0 && t >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}
