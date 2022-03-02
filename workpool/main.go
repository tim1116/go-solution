package main

import (
	"fmt"
	"log"
	"time"
	"workpool"
)

func main() {
	pool := workpool.NewPool(3)
	pool.RecoverHand = func(i interface{}) {
		log.Println(i)
	}

	start := time.Now()

	taskNum := 5
	for i := 0; i < taskNum; i++ {
		pool.Do(func(param ...interface{}) {
			fmt.Printf("task %d 开始执行\n", param[0])
			// 执行的业务代码
			time.Sleep(time.Second)
			fmt.Printf("task %d 执行完成\n", param[0])

		}, i)
	}
	pool.Wait()

	dur := time.Now().Sub(start).Seconds()
	fmt.Println("任务执行时间: ", dur)
}
