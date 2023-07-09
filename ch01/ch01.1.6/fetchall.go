package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// fetchall并发获取URL并报告它们的时间和大小
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //启动一个goroutine
	}
	for range os.Args[1:] {
		fmt.Printf(<-ch) //从通道ch接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //发送到通道ch
		return
	}
	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close() //不要泄露资源
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d  %s\n", secs, nbytes, url)
}
