package main

import "testing"

func Test416(t *testing.T) {
	canPartition([]int{66, 90, 7, 6, 32, 16, 2, 78, 69, 88, 85, 26, 3, 9, 58, 65, 30, 96, 11, 31, 99, 49, 63, 83, 79, 97, 20, 64, 81, 80, 25, 69, 9, 75, 23, 70, 26, 71, 25, 54, 1, 40, 41, 82, 32, 10, 26, 33, 50, 71, 5, 91, 59, 96, 9, 15, 46, 70, 26, 32, 49, 35, 80, 21, 34, 95, 51, 66, 17, 71, 28, 88, 46, 21, 31, 71, 42, 2, 98, 96, 40, 65, 92, 43, 68, 14, 98, 38, 13, 77, 14, 13, 60, 79, 52, 46, 9, 13, 25, 8})
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

	f := make([]int64, target+1)
	f[0] = 1
	for _, x := range nums {
		for j := target; j >= x; j-- {
			f[j] += f[j-x]
		}
	}

	if f[target] > 0 {
		return true
	}
	return false
}
