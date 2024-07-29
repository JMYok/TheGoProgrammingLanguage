package main

// 参考：https://juejin.cn/post/7393310486149316619
import (
	"fmt"
	"sync"
)

// Generator 结构体，包含一个互斥锁和一个任务ID计数器
type Generator struct {
	mutex  sync.Mutex
	taskID int
}

// NewGenerator 创建一个新的 Generator
func NewGenerator() *Generator {
	return &Generator{taskID: 0}
}

// Generate 以线程安全的方式生成一个新的唯一任务ID
func (g *Generator) Generate() int {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.taskID++
	return g.taskID
}

// worker 函数，模拟一个工作线程，接收一个生成器和一个 WaitGroup
func worker(id int, gen *Generator, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		taskID := gen.Generate()
		fmt.Printf("Worker %d 生成了任务ID: %d\n", id, taskID)
	}
}

func main() {
	gen := NewGenerator()
	var wg sync.WaitGroup

	// 启动3个工作线程，每个线程调用 worker 函数生成任务ID
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, gen, &wg)
	}

	// 使用 WaitGroup 确保所有goroutine完成后主程序才退出
	wg.Wait()
}
