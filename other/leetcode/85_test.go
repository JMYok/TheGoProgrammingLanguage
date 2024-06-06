package main

import (
	"fmt"
	"testing"
)

func Test85(t *testing.T) {
	matrix := [4][7]byte{{'0', '0', '0', '0', '0', '0', '1'}, {'0', '0', '0', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1'}, {'0', '0', '0', '1', '1', '1', '1'}}
	fmt.Println(maximalRectangle(matrix))
	fmt.Println(maximalRectangleCorrect(matrix))
}

func maximalRectangle(matrix [4][7]byte) (ans int) {
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i := range left {
		left[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j > 0 {
				left[i][j] = left[i][j-1] + 1
			} else {
				left[i][j] = 1
			}
		}
	}
	for j := 0; j < n; j++ {
		up := make([]int, m)
		down := make([]int, m)
		var st []int
		for r := 0; r < m; r++ {
			h := left[r][j]
			if len(st) > 0 && h <= left[st[len(st)-1]][j] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 {
				up[r] = st[len(st)-1]
			} else {
				up[r] = -1
			}
			st = append(st, r)
		}
		st = st[:0]
		for r := m - 1; r >= 0; r-- {
			h := left[r][j]
			if len(st) > 0 && h <= left[st[len(st)-1]][j] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 {
				down[r] = st[len(st)-1]
			} else {
				down[r] = m
			}
			st = append(st, r)
		}

		for i := 0; i < m; i++ {
			h := left[i][j]
			ans = max(ans, h*(down[i]-up[i]-1))
		}
	}
	return ans
}

func maximalRectangleCorrect(matrix [4][7]byte) (ans int) {
	n := len(matrix)
	m := len(matrix[0])
	// 各个点左边有多少个1
	left := make([][]int, n)
	for i := range left {
		left[i] = make([]int, m)
	}
	for row, rowv := range matrix {
		for col, v := range rowv {
			if v == '0' {
				continue
			}
			if col > 0 {
				left[row][col] = left[row][col-1] + 1
			} else {
				left[row][col] = 1
			}
		}
	}

	for j := 0; j < m; j++ {
		var st []int
		//上方最近最小下标位置
		up := make([]int, n)
		down := make([]int, n)
		for idx, l := range left {
			for len(st) > 0 && l[j] <= left[st[len(st)-1]][j] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 {
				up[idx] = st[len(st)-1]
			} else {
				up[idx] = -1
			}
			st = append(st, idx)
		}
		st = st[:0]
		for idx := n - 1; idx >= 0; idx-- {
			for len(st) > 0 && left[idx][j] <= left[st[len(st)-1]][j] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 {
				down[idx] = st[len(st)-1]
			} else {
				down[idx] = n
			}
			st = append(st, idx)
		}
		for i, l := range left {
			h := down[i] - up[i] - 1
			area := h * l[j]
			ans = max(ans, area)
		}
	}
	return ans
}
