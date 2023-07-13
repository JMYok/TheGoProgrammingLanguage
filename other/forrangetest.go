package main

import "fmt"

func main() {
	arr := [2]int{1, 2}
	res := []*int{}
	for _, v := range arr {
		res = append(res, &v)
	}

	fmt.Println(*res[0], *res[1])
}
