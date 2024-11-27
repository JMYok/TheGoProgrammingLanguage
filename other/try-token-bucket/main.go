package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limit := rate.NewLimiter(rate.Limit(1), 2)
	for {
		if err := limit.Wait(context.Background()); err == nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
