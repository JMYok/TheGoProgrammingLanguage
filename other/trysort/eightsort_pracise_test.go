package trysort

import (
	"fmt"
	"testing"
)

func TestQuickSortT(t *testing.T) {
	array := []int{99, 8, 7, 6, 5, 4, 32, 33}
	quickSortT(array, 0, len(array)-1)
	fmt.Println(array)
}

func quickSortT(array []int, low, high int) {
	if low >= high {
		return
	}
	pIdx := partitionT(array, low, high)
	quickSortT(array, low, pIdx-1)
	quickSortT(array, pIdx+1, high)
}

func partitionT(array []int, low, high int) int {
	left, right := low, high
	pivot := array[low]
	idx := low
	for left <= right {
		for left <= right {
			if array[right] < pivot {
				array[right], array[idx] = array[idx], array[right]
				idx = right
				right--
			}
			right--
		}
		for left <= right {
			if array[left] >= pivot {
				array[left], array[idx] = array[idx], array[left]
				idx = left
				left++
			}
			left++
		}
	}
	array[idx] = pivot
	return idx
}

func heapSortT(array []int) {
	//
}

func createHeapT(array []int) {

}

func downAdjustT(array []int, low, high int) {

}

func bubbleSortT(array []int) []int {

	return nil
}
