package huawei

import (
	"testing"
)

var matrix [15][15]bool
var m, n int
var ans = int(1e9)

func check(x, y, length int) bool {
	// 判断从第x行第y个格子开始能否填充边长为length的正方形
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if matrix[x+i][y+j] {
				return false
			}
		}
	}
	return true
}

func fill(x, y, length int) {
	// 采用异或方法进行填充/取消填充
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[x+i][y+j] = !matrix[x+i][y+j]
		}
	}
}

func dfs(x, y, cnt int) {
	if cnt >= ans {
		return // 剪枝，当前记录数已经≥已记录的最小答案，不再进行搜索
	}
	if x == m {
		ans = cnt // 填充完毕，更新答案
		return
	}
	if y >= n {
		dfs(x+1, 0, cnt)
	}
	full := true
	for i := y; i < n; i++ {
		// 从当前行的第y个格子开始枚举，找到第一个没有填充的格子
		if !matrix[x][i] {
			// 当前格子未填充，尝试填充正方形
			full = false
			for j := min(n-i, m-x); j >= 0; j-- {
				// 枚举填充正方形的边长，从长边开始枚举
				if check(x, i, j) {
					fill(x, i, j)      // 填充
					dfs(x, y+j, cnt+1) // 填充完一个正方形，尝试下一次填充
					fill(x, i, j)      // 取消填充
				}
			}
			break // 尝试在当前格子填充正方形的所有情况已经全部考虑，直接跳出循环
		}
	}
	if full {
		dfs(x+1, 0, cnt) // 当前行都填充了，搜索下一行
	}
}

func TestDfs(t *testing.T) {
	m, n = 3, 2
	dfs(0, 0, 0)
	t.Log(ans)
}
