package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. 声明 sync.WaitGroup 类型的变量 wg
	var wg sync.WaitGroup
	fmt.Printf("Program start \n")
	for i := 0; i < 5; i++ {
		// 2. 每启动一个 goroutine，则调用 wg.Add(1)
		wg.Add(1)
		go concurrentTaks(&wg, i)
	}
	// 4. 等待所有的 goroutine 执行完成
	wg.Wait()
	fmt.Printf("Program end \n")
}

func concurrentTaks(wg *sync.WaitGroup, taskNumber int) {
	// 3. 函数执行完成时调用 wg.Done
	defer wg.Done()
	fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d \n", taskNumber)
}
