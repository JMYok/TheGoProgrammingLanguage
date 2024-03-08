package slice

import (
	"testing"
)

func TestSliceCap(t *testing.T) {
	slice1, slice2 := make([]int, 0, 1), make([]int, 0, 10)
	//期待容量大于原来容量的2倍
	slice1 = append(slice1, []int{1, 2, 3, 4, 5}...)
	t.Log(len(slice1), cap(slice1))
	//期待容量小于原来容量的2倍(扩容前容量小于256)
	slice2 = append(slice2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}...)
	t.Log(len(slice2), cap(slice2))
	//期待容量小于原来容量的2倍(扩容前容量大于256)

}
