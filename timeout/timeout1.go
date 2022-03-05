package main

import (
	"fmt"
	"time"
)

var chanTimeOut = make(chan bool, 1)
var data = make(chan string, 1)

func main() {
	timeout := 3 * time.Second
	go func(ch chan<- bool) {
		time.Sleep(timeout)
		ch <- true
	}(chanTimeOut)

	go do()

	select {
	case <-chanTimeOut:
		fmt.Println("超时~")
	case rev := <-data:
		fmt.Println("收到数据:", rev)
	}

	fmt.Println("end")
}

// 执行某项业务逻辑
func do() {
	time.Sleep(10 * time.Second)
	data <- "业务执行完成"
	close(data)
}
