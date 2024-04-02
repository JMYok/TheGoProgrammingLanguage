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
