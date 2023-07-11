package main

import "fmt"

func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("出现了panic,使用了recover获取信息：", error)
		}
	}()
	fmt.Println("11111111")
	panic("出现panic")
	fmt.Println("22222222")
}
