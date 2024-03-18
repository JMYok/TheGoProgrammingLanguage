package main

import (
	"fmt"
	"unicode"
)

func main() {
	cntLower := 0
	cntUpper := 0
	s := "Hello World"
	for _, c := range s {
		if unicode.IsLower(c) {
			cntLower++
		} else if unicode.IsUpper(c) {
			cntUpper++
		}
	}
	fmt.Println(cntLower)
	fmt.Println(cntUpper)
}
