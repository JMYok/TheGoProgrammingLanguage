package main

import "fmt"

func main() {
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
