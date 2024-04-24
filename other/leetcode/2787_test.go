package main

import "testing"

func Test2787(t *testing.T) {
	numberOfWays(218, 1)
}

func numberOfWays(n int, x int) int {
	// f[i+1][j] = f[i][j]+f[i][j-x]
	nums := make([]int, 0)
	p := (int)(1e9 + 7)
	for num := 1; qmi(num, x, p) <= n; num++ {
		nums = append(nums, qmi(num, x, p))
	}
	f := make([]int, n+1)
	f[0] = 1
	for _, x := range nums {
		for j := n; j >= x; j-- {
			f[j] += f[j-x]
		}
	}
	return f[n]
}

func qmi(m, k, p int) int {
	res := 1 % p
	t := m
	for k > 0 {
		if k&1 == 1 {
			res = res * t % p
		}
		t = t * t % p
		k >>= 1
	}
	return res
}
