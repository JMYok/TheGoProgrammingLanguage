package binsearch

type BinSearch struct {
	arr []int
}

func NewBinSearch(arr []int) *BinSearch {
	return &BinSearch{
		arr: arr,
	}
}

func (b *BinSearch) BinarySearchVal(x int) int {
	l, r := 0, len(b.arr)-1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if b.arr[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return b.arr[l]
}
