package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var data = make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go func() {
		time.Sleep(10 * time.Second)
		data <- "执行业务完成"
		close(data)
	}()

	select {
	case rec := <-data:
		fmt.Println("receive:", rec)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	fmt.Println("end")
}
