package main

import (
	"fmt"
	"os"
)

func main() {
	for index, val := range os.Args[1:] {
		fmt.Println(index, val)
	}
}
