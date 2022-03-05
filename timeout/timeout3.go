package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 控制多个任务
	go work1(ctx)
	go work2(ctx)

	// 此处控制超时
	time.Sleep(time.Second * 3)
	cancel()

	time.Sleep(time.Second * 2)
	fmt.Println("end")
}

func work1(ctx context.Context) {
	ctx2, _ := context.WithCancel(ctx)

	go work3(ctx2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work2 超时~")
			return
		case <-time.After(time.Second):
			// block
		}
	}
}

func work3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work3 超时~")
			return
		case <-time.After(time.Second):
			// block
		}
	}
}

func work2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work1 超时~")
			return
		case <-time.After(time.Second):
			// block
		}
	}
}
