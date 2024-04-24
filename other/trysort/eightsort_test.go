package trysort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println(bubbleSort([]int{3, 2, 1}))
}

func TestHeapSort(t *testing.T) {
	array := make([]int, 7)
	array[1] = 8
	array[2] = 6
	array[3] = 5
	array[4] = 3
	array[5] = 1
	array[6] = 2
	fmt.Println(heapSort(array))
}

func TestQuickSort(t *testing.T) {
	array := []int{99, 8, 7, 6, 5, 4, 32, 33}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

// **************** 快速排序 从小到大***********
func quickSort(array []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIdx := partition(array, startIndex, endIndex)
	quickSort(array, startIndex, pivotIdx-1)
	quickSort(array, pivotIdx+1, endIndex)
}

func partition(array []int, startIndex int, endIndex int) int {
	// 选第一个数为主元
	pivot := array[startIndex]
	//记录当前没有数字的坑位，也是最终主元的位置
	index := startIndex

	left, right := startIndex, endIndex

	for left <= right {
		//从右到左找小于pivot的数
		for left <= right {
			if array[right] < pivot {
				array[left], array[right] = array[right], array[left]
				index = right
				right--
				break
			}
			right--
		}

		//从左到右找大于pivot的数
		for left <= right {
			if array[left] >= pivot {
				array[left], array[right] = array[right], array[left]
				index = left
				left++
				break
			}
			left++
		}
	}
	array[index] = pivot
	return index
}

// **************** 堆排序 从大到小****************
func heapSort(array []int) []int {
	//小顶堆
	createMinHeap(array)
	for i := len(array) - 1; i > 1; i-- {
		array[1], array[i] = array[i], array[1]
		downAdjust(1, i-1, array)
	}
	return array
}

func createMinHeap(array []int) {
	for i := len(array) / 2; i >= 1; i-- {
		downAdjust(i, len(array)-1, array)
	}
}

// 向下调整(小顶堆)
func downAdjust(low, high int, heap []int) {
	i := low
	j := 2 * i
	for j <= high {
		if j+1 <= high && heap[j+1] < heap[j] {
			j++
		}
		if heap[i] > heap[j] {
			heap[i], heap[j] = heap[j], heap[i]
			i = j
			j = 2 * i
		} else {
			break
		}
	}
}

// ********************冒泡排序******************
func bubbleSort(array []int) []int {
	n := len(array)
	sortBorder := n - 1
	lastIdx := n - 1
	for range array {
		for j := 0; j < sortBorder; j++ {
			isSorted := true
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSorted = false
				lastIdx = j
			}
			if isSorted {
				break
			}
		}
		sortBorder = lastIdx
	}
	return array
}
