package main

import (
	"bufio"
	"fmt"
	"os"
)

// dup1输出标准输入中出现次数大于1的行，前面是次数

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count=%d\tval=%s\n", n, line)
		}
	}
}
