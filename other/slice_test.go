package main

import (
	"fmt"
	"testing"
)

func TestMockStack(t *testing.T) {
	const N = 10
	tt := 0 //指向栈顶，不是栈顶下一个数；tt表示栈的大小
	stack := make([]int, N)

	// 入栈
	tt++
	stack[tt] = 66
	tt++
	stack[tt] = 88
	fmt.Println(stack)

	//出栈
	tt--
	fmt.Println(stack[tt])

	//不为空
	if tt > 0 {
		fmt.Println("stack不为空")
	}
}

func TestSlice(t *testing.T) {
	slice := make([]int, 3, 4)
	fmt.Println(slice)

	arr1 := make([]int, 0, 4)
	arr1 = append(arr1, 1)
	arr2 := append(arr1, 2)
	arr3 := append(arr1, 3)
	fmt.Printf("arr1=%v,addr1=%p\n", arr1, &arr1)
	fmt.Printf("arr2=%v,addr2=%p\n", arr2, &arr2)
	fmt.Printf("arr3=%v,addr3=%p\n", arr3, &arr3)

	arr := make([]int, 4, 4)
	fmt.Printf("cap=%d\n", cap(arr))
	arr = append(arr, 1)
	fmt.Printf("cap=%d\n", cap(arr))
}
