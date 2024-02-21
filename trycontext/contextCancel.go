package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go watch(ctx, "gorountine1")
	go watch(ctx, "gorountine2")

	time.Sleep(6 * time.Second)
	fmt.Println("end watching")
	cancelFunc()
	time.Sleep(time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("over:", name) //调用cancel之后,Done会收到信号
			return
		default:
			fmt.Println("watching:", name)
			time.Sleep(time.Second)
		}
	}
}
