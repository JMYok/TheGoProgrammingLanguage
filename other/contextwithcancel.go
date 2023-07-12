package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//派生出带有返回函数cancel的ctx
	ctx, cancel := context.WithCancel(context.Background())
	go Watch(ctx, "goroutine1")
	go Watch(ctx, "goroutine2")

	time.Sleep(6 * time.Second)
	fmt.Println("end watching!")
	cancel() //通知goroutine1和goroutine2关闭
	time.Sleep(1 * time.Second)
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s exit!\n", name)
			return
		default:
			fmt.Printf("%s watching...\n", name)
			time.Sleep(time.Second)
		}
	}
}
