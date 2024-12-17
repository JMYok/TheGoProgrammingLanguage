package binsearch

import "testing"

func TestFunc(t *testing.T) {
	search := NewBinSearch([]int{1, 2, 3, 4, 5, 6})
	res := search.BinarySearchVal(3)
	if res == 3 {
		t.Log("success")
	} else {
		t.Errorf("expected %v,but got %v", 3, res)
	}
}
