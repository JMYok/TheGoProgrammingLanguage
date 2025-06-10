package main

import (
	"fmt"
	"iter" // 导入 iter 包
)

// Backward 返回一个逆序遍历切片的迭代器函数
//func Backward[E any](s []E) iter.Seq2[int, E] {
//	// 返回符合 iter.Seq2 签名的函数
//	return func(yield func(int, E) bool) {
//		// 从后往前遍历
//		for i := len(s) - 1; i >= 0; i-- {
//			// 调用 yield "产出" 索引和元素
//			// 如果 yield 返回 false，立即停止迭代
//			if !yield(i, s[i]) {
//				return
//			}
//		}
//	}
//}

func BackwardV[E any](s []E) iter.Seq[int] {
	// 返回符合 iter.Seq2 签名的函数
	return func(yield func(int) bool) {
		// 从后往前遍历
		for i := len(s) - 1; i >= 0; i-- {
			// 调用 yield "产出" 索引和元素
			// 如果 yield 返回 false，立即停止迭代
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	data := []string{"a", "b", "c"}

	// 使用 for range 遍历我们自定义的 Backward 迭代器
	//for i, s := range Backward(data) {
	//	fmt.Printf("Index: %d, Value: %s\n", i, s)
	//	if i == 1 { // 演示提前 break
	//		break
	//	}
	//}

	// 将 Push 迭代器 s 转换为 Pull 迭代器
	// next() 返回下一个元素和是否有效的布尔值
	// stop() 用于提前释放迭代器可能持有的资源 (如果有的话)
	next, stop := iter.Pull(BackwardV(data))
	defer stop() // 推荐使用 defer stop()

	for {
		value, ok := next()
		if !ok { // 没有更多元素了
			break
		}
		fmt.Printf("PULL:%d\n", value)
	}
}
