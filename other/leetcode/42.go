package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}

func trap(height []int) int {
	var st []int
	var ans int
	for i, x := range height {
		for len(st) > 0 && x > height[st[len(st)-1]] {
			if len(st) >= 2 {
				h := min(height[st[len(st)-2]], x) - height[st[len(st)-1]]
				w := i - st[len(st)-2] - 1
				ans += h * w
			}
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}

func Min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
