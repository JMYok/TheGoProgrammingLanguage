package main

func main() {
	/*未初始化channel*/
	//var ch chan int
	//panic: close of nil channel
	//close(ch)

	//接收 fatal error: all goroutines are asleep - deadlock!
	//ch <- 1

	//发送 fatal error: all goroutines are asleep - deadlock!
	//x := <-ch
	//println(x)

	/*已关闭的channel*/
	ch := make(chan int)
	ch <- 1

	//panic: close of closed channel
	close(ch)
	//close(ch)

	//ch <- 2
	x := <-ch
	//y := <-ch
	println(x)
	//println(y)

}
