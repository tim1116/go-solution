package main

import (
	"fmt"
	"time"
)

func main() {
	var data = make(chan string, 1)
	go func() {
		time.Sleep(10 * time.Second)
		data <- "业务执行完成"
		close(data)
	}()

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("超时~")
	case rev := <-data:
		fmt.Println("收到数据:", rev)
	}

	fmt.Println("end")
}
