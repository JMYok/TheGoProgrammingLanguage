package main

import (
	"fmt"
	"sync"
)

func withGoroutine(opts ...func() error) (err error) {
	var wg sync.WaitGroup
	for _, opt := range opts {
		wg.Add(1)
		go func(handler func() error) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("recover:%v\n", e)
				}
				wg.Done()
			}()
			//真正的逻辑调用
			e := handler()
			if err == nil && e != nil {
				err = e
			}
		}(opt)
	}

	//等待所有协程执行完
	wg.Wait()

	return err
}

func main() {
	handler1 := func() error {
		panic("handler1 fail")
		return nil
	}

	handler2 := func() error {
		panic("handler2 fail")
		return nil
	}

	handler3 := func() error {
		panic("handler3 fail")
		return nil
	}
	handler4 := func() error {
		panic("handler4 fail")
		return nil
	}
	handler5 := func() error {
		panic("handler5 fail")
		return nil
	}

	err := withGoroutine(handler1, handler2, handler3, handler4, handler5)

	if err != nil {
		fmt.Printf("err is %v\n", err)
	}
}
