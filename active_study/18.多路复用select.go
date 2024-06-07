package main

import (
	"fmt"
	"time"
)

func main() {
	// ch1 := make(chan int, 2)
	// ch2 := make(chan int, 2)
	// ch2 <- 5
	// select {
	// // 1. 向 ch1 写入数据
	// case ch1 <- 10:
	// 	fmt.Println("send to ch1")
	// // 2. 从 ch2读取数据
	// case recv := <-ch2:
	// 	fmt.Println("receive from ch2", recv)
	// // 3. 默认行为，当每一个 case 都不能执行时会执行 default 里面的逻辑
	// default:
	// 	fmt.Println("running default")
	// }

	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("slept for 3 seconds")
		// 等待三秒后 关闭 channel
		close(ch1)
	}()

	//
	select {
	case <-ch1:
		fmt.Println("finished")
	}
}
