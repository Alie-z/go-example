package main

import (
	"fmt"
	"math/rand"
	"time"
)

// work  业务逻辑执行函数，基于随机数模拟每次执行的不同耗时
func work(i int, ch chan<- int) {
	// 1. 随机一个[0,5)之间的值，模拟程序处理耗时
	rand.Seed(time.Now().UnixNano())
	dura := rand.Intn(5)
	fmt.Printf("goroutine %d sleep dura is set to %d seconds\n", i, dura)
	time.Sleep(time.Duration(dura) * time.Second)
	ch <- i
}

func main() {
	ch1 := make(chan int)
	// 2. 启动三个 goroutine 并发执行业务逻辑
	go work(1, ch1)
	go work(2, ch1)
	go work(3, ch1)

	select {
	case recv := <-ch1:
		fmt.Println("finished, receive ret from goroutine ", recv)
	// 2. 基于 time.Second设置超时两秒钟
	case <-time.After(2 * time.Second):
		fmt.Println("timeout after 2 second")
	}
}
