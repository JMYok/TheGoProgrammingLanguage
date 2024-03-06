package main

import "fmt"

func main() {
	s := make([]int, 0, 100)
	//PrintSliceStruct(&s)
	//test(s)
	testAppend(&s)
	fmt.Println(s)
}

func testAppend(s *[]int) {
	*s = append(*s, 666)
	fmt.Println("testAppend:", s)
}

//func test(s []int) {
//	PrintSliceStruct(&s)
//}
//
//func PrintSliceStruct(s *[]int) {
//	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
//
//	fmt.Printf("slice struct:%+v, slice is %v\n", ss, s)
//}
