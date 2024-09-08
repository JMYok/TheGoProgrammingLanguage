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
	l, r := low, high
	pivot := array[low]
	idx := low
	for l <= r {
		for l <= r {
			if array[r] < array[idx] {
				array[r], array[idx] = array[idx], array[r]
				idx = r
				r--
				break
			}
			r--
		}
		for l <= r {
			if array[l] >= array[idx] {
				array[l], array[idx] = array[idx], array[l]
				idx = l
				l++
				break
			}
			l++
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
