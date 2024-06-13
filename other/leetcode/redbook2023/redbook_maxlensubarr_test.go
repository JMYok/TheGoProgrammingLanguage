package main

import (
	"fmt"
	"testing"
)

func TestProc(test *testing.T) {
	var t int
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		var n, x int
		fmt.Scanf("%d %d\n", &n, &x)
		preSum := make([]int, n+1)
		for i := 0; i < n; i++ {
			var num int
			if i < n-1 {
				fmt.Scanf("%d ", &num)
			} else {
				fmt.Scanf("%d\n", &num)
			}
			preSum[i+1] = preSum[i] + num
		}

		ans := -int(1e9 + 5)
		minV := preSum[0]
		for i := 1; i <= n; i++ {
			ans = max(ans, preSum[i]-minV)
			if minV > preSum[i] {
				minV = preSum[i]
			}
			if minV > x {
				minV = x
			}
		}
		fmt.Println(ans)
		t--
	}
}
