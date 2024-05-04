package main

import (
	"math"
	"testing"
)

func Test84(t *testing.T) {
	largestRectangleArea([]int{2, 2, 2, 2})
}

func largestRectangleArea(heights []int) int {
	// 2 1 5 6 2 3 0
	// st=[1 ]
	// ans=[6 2 10 6 8 3]
	// 2 2 2
	// st[2]
	// ans=[2,4,6]
	// 2 1 2 0
	// st [1 ]
	// ans= [2,2,2,0]

	// j>i,h[j]<=h[i]时：
	//1. 计算栈中元素i的右边矩形（不包含h[i]）:ans[i]+= (st[j]-st[i]-1)*h[i],弹栈
	// 2.计算j左边矩形（不包含h[j]），ans[j]+=弹栈个数*h[j]
	// j>i,h[j]>h[i]时：入栈
	var st []int
	heights = append(heights, 0)
	ans := make([]int, len(heights))
	copy(ans, heights)
	for i, x := range heights {
		popNum := 0
		for len(st) > 0 && x <= heights[st[len(st)-1]] {
			// 计算右边
			topIdx := st[len(st)-1]
			ans[topIdx] += (i - topIdx - 1) * heights[topIdx]
			st = st[:len(st)-1]
			popNum++
		}
		// 计算左边
		ans[i] += popNum * heights[i]
		st = append(st, i)
	}
	res := math.MinInt
	for _, x := range ans {
		res = max(res, x)
	}
	return res
}
