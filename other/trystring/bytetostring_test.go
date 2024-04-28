package main

import (
	"fmt"
	"testing"
)

func TestByteString(t *testing.T) {
	var str string = "test"
	var data []byte = []byte(str)
	fmt.Println(data)
}

func TestStringToByte(t *testing.T) {
	var str string = "ab"
	fmt.Println(str[0])
	fmt.Printf("%c\n", str[0])
}
