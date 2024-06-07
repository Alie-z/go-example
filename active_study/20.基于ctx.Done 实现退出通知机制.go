package main

import (
	"context"
	"fmt"
	"time"
)

func work1(ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("G2 finished")
			return
		// 这里利用 default 不阻塞的原理， 实现任务处理
		default:
			time.Sleep(time.Second)
			fmt.Println("G2 is still working at ", i)
		}
	}
}

func main() {
	// 主 goroutine G0
	fmt.Println("G0 running")
	ctx := context.Background()
	//1. G0 启动子 goroutine G1
	go func(ctx context.Context) {
		// 2. 一个带 cancel 方法的新的 context
		ctx, cancel := context.WithCancel(ctx)
		// 3. G1 启动子 goroutine G2
		go work1(ctx)
		time.Sleep(5 * time.Second)
		// 4. 调用 cancel 通知G2结束
		cancel()
		fmt.Println("G1 finished")
	}(ctx)
	var ch chan int
	// 3. 永久阻塞G0，可以 CTRL+C 结束执行
	<-ch
}
