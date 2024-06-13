package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var a [3]int
	//var b [3]int = [3]int{1, 2, 3}
	//fmt.Printf("type of b is: %v, value of b is :%v\n", reflect.TypeOf(b), reflect.ValueOf(b))

	//var nullList []int
	//testNullArrayType(&nullList)
	//
	//fmt.Printf("value of nullList is: %v\n", nullList)

	//var c = [3]int{4, 5, 6}
	//for i, val := range a {
	//	fmt.Printf("%d:%d\n", i, val)
	//}
	//for i, val := range b {
	//	fmt.Printf("%d:%d\n", i, val)
	//}
	//for i, val := range c {
	//	fmt.Printf("%d:%d\n", i, val)
	//}

	a := []int{0, 0, 0}
	b := []int{1, 1, 1}
	fmt.Println("a:", a, "b:", b)
	switchArray(a, b)
	fmt.Println("a:", a, "b:", b)
}

func switchArray(a, b []int) {
	var c []int
	c = a
	a = b
	b = c
}

func testNullArrayType(modelList interface{}) {
	v := reflect.ValueOf(modelList)
	fmt.Printf("value of modelList is: %v\n", v) //value is an address pointed to the []int
	t := reflect.Indirect(v).Type()
	fmt.Printf("type of modelList is: %v\n", t)
}
