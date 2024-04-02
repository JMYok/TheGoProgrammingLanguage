package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}

// 单调栈
type Stack struct {
	data []int
	len  int
}

func (s *Stack) push(x int) {
	s.data = append(s.data, x)
	s.len++
}

func (s *Stack) pop() int {
	val := s.data[s.len-1]
	s.data = s.data[:s.len-1]
	s.len--
	return val
}

func (s *Stack) peek() int {
	return s.data[s.len-1]
}

func (s *Stack) empty() bool {
	return s.len == 0
}

func trap(height []int) int {
	st := new(Stack)
	st.data = make([]int, 0)
	ans := 0
	i := 0
	for i < len(height) {
		for !st.empty() && height[i] > height[st.peek()] {
			top := st.pop()
			//说明没有左边界了
			if st.empty() {
				break
			}
			//左边界
			left := st.peek()
			width := i - left - 1
			h := Min(height[left], height[i]) - height[top]
			ans += width * h
		}
		st.push(i)
		i++
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

// //dp方法
// func trap(height []int) int {
// 	length := len(height)
// 	//记录第i个节点左边最高的墙
// 	leftMax := 0
// 	//记录第i个节点右边最高的墙
// 	rightMax := 0
// 	ans := 0
// 	//计算当前位置能收集多少雨水
// 	left, right := 1, length-2
// 	for left <= right {
// 		//从左向右
// 		if height[left-1] < height[right+1] {
// 			leftMax = Max(height[left-1], leftMax)
// 			if height[left] < leftMax {
// 				ans += leftMax - height[left]
// 			}
// 			left++
// 		} else {
// 			//从右向左
// 			rightMax = Max(height[right+1], rightMax)
// 			if height[right] < rightMax {
// 				ans += rightMax - height[right]
// 			}
// 			right--
// 		}
// 	}
// 	return ans
// }

// func Max(x, y int) int {
// 	if x >= y {
// 		return x
// 	} else {
// 		return y
// 	}
// }

// func Min(x, y int) int {
// 	if x <= y {
// 		return x
// 	} else {
// 		return y
// 	}
// }
