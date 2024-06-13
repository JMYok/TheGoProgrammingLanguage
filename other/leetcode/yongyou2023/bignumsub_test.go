package main

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	// var stra, strb string
	// fmt.Scan(&stra, &strb)
	stra := "1"
	strb := "2"
	i := len(stra) - 1
	j := len(strb) - 1
	ans := ""
	var t int
	for i >= 0 && j >= 0 {
		t = int(stra[i]) - int(strb[j]) + t
		if t < 0 && i > 0 {
			ans += string(byte(t) + 10)
			t = -1
		} else {
			ans += string(t)
			t = 0
		}
		i--
		j--
	}
	for i >= 0 {
		ans += string(stra[i] + byte(t))
		t = 0
		i--
	}
	for j >= 0 {
		ans += string(strb[j] + byte(t))
		t = 0
		j--
	}
	fmt.Println(ans)
}
