// // 有缓冲

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// 1. 构造有缓冲区 channel
// 	ch1 := make(chan int, 1)
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("slept 1 second")
// 		i := <-ch1
// 		fmt.Println("received", i)
// 	}()
// 	// 2. 向 ch 中写入数字 0
// 	ch1 <- 9
// 	// 3. 写入 channel 成功，打印字符串
// 	fmt.Println("written success")
// 	// 4. sleep 2秒等待所有的协程执行完成
// 	time.Sleep(2 * time.Second)
// }

// // 无缓冲

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// 1. 构造无缓冲区 channel
// 	ch1 := make(chan int)
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("slept 1 second")
// 		i := <-ch1
// 		fmt.Println("received", i)
// 	}()
// 	// 2. 向 ch 中写入数字 0
// 	ch1 <- 8
// 	// 3. 写入 channel 成功，打印字符串
// 	fmt.Println("written success")
// 	// 4. sleep 2秒等待所有的协程执行完成
// 	time.Sleep(2 * time.Second)
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	syncCh := make(chan bool)
	// 1. 运行一个 goroutine
	go func() {
		longTask2()
		fmt.Println("task2 finished")
		syncCh <- true
	}()
	longTask()
	fmt.Println("task1 finished")
	// 2. 会一直阻塞直到 longTask2执行完
	// 如果 注释掉这一行，则输出里只能看到 task1 finished
	<-syncCh
}

func longTask2() {
	time.Sleep(3 * time.Second)
}

func longTask() {
	time.Sleep(1 * time.Second)
}
