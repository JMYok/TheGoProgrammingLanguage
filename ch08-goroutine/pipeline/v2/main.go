package main

import "fmt"

func main() {
	//generator、multiply、add代表管道的不同阶段。每个阶段会返回一个通道供下一个阶段消费
	//通道 done 则是为了实现协程的退出而设计的
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		// 非缓冲通道
		multipliedStream := make(chan int)
		// 协程仅在 done 通道被关闭或者 intStream 被全部读取完时才会退出
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		// 调用协程后，立即返回 multipliedStream 通道
		// 主协程调用 add 后会立刻得到返回的 multipliedStream，并继续把这个通道传给下游的 add 函数。
		// 整个流水线结构的执行通过管道将数据逐步传递下去，而无需等待协程完成
		return multipliedStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			// 如果没有值会阻塞在这里
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	// 生成数字流
	intStream := generator(done, 1, 2, 3, 4)
	// 执行流
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
