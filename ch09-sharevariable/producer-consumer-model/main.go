package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 模拟订单对象
type OrderInfo struct {
	id int
}

// 生产订单--生产者
func producerOrder(out chan<- OrderInfo) {
	// 业务生成订单
	for i := 0; i < 10; i++ {
		order := OrderInfo{id: i + 1}
		fmt.Println("生成订单，订单ID为：", order.id)
		out <- order // 写入channel
	}
	// 如果不关闭，消费者就会一直阻塞，等待读
	close(out) // 订单生成完毕，关闭channel
	wg.Done()
}

// 处理订单--消费者
func consumerOrder(in <-chan OrderInfo, idx int) {
	// 从channel读取订单，并处理
	for order := range in {
		fmt.Printf("%d号消费者读取订单，订单ID为：%d\n", idx, order.id)
	}
	wg.Done()
}

func main() {
	ch := make(chan OrderInfo, 5)
	wg.Add(2)
	go producerOrder(ch)
	go consumerOrder(ch, 1)
	go consumerOrder(ch, 2)
	go consumerOrder(ch, 3)
	wg.Wait()
}
