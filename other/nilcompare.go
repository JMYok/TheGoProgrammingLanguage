package main

import "fmt"

func main() {
	var p *int = nil
	var i interface{} = p
	fmt.Println(i == p)
	fmt.Println(p == nil)

	//这里会将nil装换为type=nil,val=nil的接口值
	//但是i为type=*int,val=nil的接口值，所以结果为false
	fmt.Println(i == nil)
}
