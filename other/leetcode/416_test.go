package main

import (
	"fmt"
	"testing"
)

func Test416(t *testing.T) {
	fmt.Println(canPartition([]int{1, 5, 10, 6}))
}

func canPartition(nums []int) bool {
	target := 0
	for _, x := range nums {
		target += x
	}
	if target%2 == 1 {
		return false
	}
	target /= 2

	n := len(nums)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, target+1)
		f[i][0] = true
	}

	for i, x := range nums {
		for j := x; j <= target; j++ {
			f[i+1][j] = f[i][j] || f[i][j-x]
		}
	}
	return f[n][target]
}
