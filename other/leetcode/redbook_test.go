package main

import (
	"fmt"
	"slices"
	"testing"
)

// 0 1 2 3 4
type pair struct {
	left  int
	right int
}

func TestRedBook(t *testing.T) {
	var n, m, k int
	fmt.Scanf("%d%d%d", &n, &m, &k)
	intervals := make([]pair, m)
	for m > 0 {
		var l, r int
		fmt.Scanf("%d%d", &l, &r)
		intervals = append(intervals, pair{l, r})
		m--
	}
	slices.SortFunc(intervals, func(a, b pair) int {
		// 先按left字段升序排序
		if a.left != b.left {
			return a.left - b.left
		}
		// 如果left字段相等，则按right字段升序排序
		return a.right - b.right
	})
	ans := 0
	for _, interval := range intervals {
		startL := interval.left
		endR := startL + k + 1
		res := 0
		var r int
		for _, inter := range intervals {
			l := inter.left
			r = inter.right
			if r > endR {
				break
			}
			res += r - l
		}
		res += endR - r
		ans = max(ans, res)
	}
	fmt.Println(ans)
}
