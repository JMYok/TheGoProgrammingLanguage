package slice

import (
	"fmt"
	"testing"
)

// 扩容测试
func TestSliceCap(t *testing.T) {
	slice1, slice2 := make([]int, 0, 1), make([]int, 0, 10)
	//期待容量大于原来容量的2倍，扩容后容量=期待容量
	slice1 = append(slice1, []int{1, 2, 3, 4, 5}...)
	t.Log(len(slice1), cap(slice1))
	//期待容量小于原来容量的2倍(扩容前容量小于256)，扩容后容量=2*原来容量
	slice2 = append(slice2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}...)
	t.Log(len(slice2), cap(slice2))
	//期待容量小于原来容量的2倍(扩容前容量大于256)

}

// 百度面试真题
func Test2(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{}
	c := []int{1, 2, 3, 4, 5}
	var d []int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		b[i] = v
	}
	//a=[1,12,13,4,5],b=[1,2,3,4,5]
	fmt.Printf("a=%v,b=%v\n", a, b)

	for i, v := range c {
		if i == 0 {
			c[1] = 12
			c[2] = 13
		}
		d = append(d, v)
	}
	//c=[1,12,13,4,5],d=[1,12,13,4,5]
	fmt.Printf("c=%v, d=%v\n", c, d)
}

// 美团面试真题
func Test1(t *testing.T) {
	doAppend := func(s []int) {
		s = append(s, 1)
		printLengthAndCapability(s)
	}
	s := make([]int, 8, 8)
	doAppend(s[:4])             //len 5 [0,0,0,0,1,0,0,0] cap 8
	printLengthAndCapability(s) //len8 len 8
	doAppend(s)                 //触发扩容 len 9 [0,0,0,0,1,0,0,0,1] cap 16
	printLengthAndCapability(s) //len 8 cap 8
}

func printLengthAndCapability(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
