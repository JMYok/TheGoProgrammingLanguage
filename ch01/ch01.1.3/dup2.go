package main

import (
	"bufio"
	"fmt"
	"os"
)

// dup2打印输入中 多次出现的行的个数和文本
// 它从stdin或指定的文件列表读取
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//若没指定文件，输入文本
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%s\n", os.Args[0])
		if n >= 1 {
			fmt.Printf("num:%d\tval:%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
