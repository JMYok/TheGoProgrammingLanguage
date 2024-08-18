package main

import "fmt"

func numberOfPatrolBlocks(block_size_row int, block_size_cols int, block [][]int32) (res int32) {
	i, j := 0, -1
	res = 0

	for {
		// 无人机是否能够运动
		flag := false

		// 向右
		for j+1 < block_size_cols && block[i][j+1] == 0 {
			j++
			res++
			block[i][j] = -1
			flag = true
		}

		// 向下
		for i+1 < block_size_row && block[i+1][j] == 0 {
			i++
			res++
			block[i][j] = -1
			flag = true
		}

		// 向左
		for j-1 >= 0 && block[i][j-1] == 0 {
			j--
			res++
			block[i][j] = -1
			flag = true
		}

		// 向上
		for i-1 >= 0 && block[i-1][j] == 0 {
			i--
			res++
			block[i][j] = -1
			flag = true
		}

		if !flag {
			break
		}
	}

	return
}

func main() {
	var res int32

	block_rows := 0
	block_cols := 0
	fmt.Scanf("%d", &block_rows)
	fmt.Scanf("%d", &block_cols)

	var block [][]int32
	for block_i := 0; block_i < block_rows; block_i++ {
		var block_row []int32
		for block_j := 0; block_j < block_cols; block_j++ {
			var block_item int32
			fmt.Scanf("%d", &block_item)
			block_row = append(block_row, block_item)
		}
		block = append(block, block_row)
	}
	fmt.Println(block)

	res = numberOfPatrolBlocks(block_rows, block_cols, block)
	fmt.Printf("%d\n", res)
}
