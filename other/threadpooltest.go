package main

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(funcArg func() error) *Task {
	return &Task{
		f: funcArg,
	}
}

type Pool struct {
	Workernum int
	Jobch     chan *Task
}

func NewPool(workerNum int, taskNum int) *Pool {
	return &Pool{
		Workernum: workerNum,
		Jobch:     make(chan *Task, taskNum),
	}
}

func (p *Pool) Size() int {
	return p.Workernum
}

// 协程池里干活的goroutine
func (p *Pool) worker(i int) {
	for task := range p.Jobch {
		if err := task.f(); err != nil {
			fmt.Printf("worker %d handle fail: %v\n", i, err)
			return
		}
		fmt.Printf("worker %d handle success\n", i)
	}
}

func (p *Pool) AddTask(task *Task) {
	p.Jobch <- task
}

func (p *Pool) Run() {
	for i := 0; i < p.Size(); i++ {
		go p.worker(i)
	}
}

func main() {
	p := NewPool(3, 10)
	p.Run()
	task := NewTask(func() error {
		fmt.Printf("I am Task\n")
		return nil
	})
	for i := 0; i < 8; i++ {
		p.AddTask(task)
	}
	time.Sleep(3 * time.Second)
}
