// package main

// import (
// 	"fmt"
// 	"time"
// )

// // ticking
// func ticking() {
// 	var times int
// 	// 构建一个无限循环
// 	for {
// 		times++
// 		fmt.Println("ticking", times)
// 		// 延时
// 		time.Sleep(2 * time.Second)
// 	}

// }

// func main() {

// 	// 1. 并发执行ticking
// 	go ticking()
// 	go ticking()

// 	// 2. 接受命令行输入, 打印输出后退出
// 	var input string
// 	fmt.Println("随意输入一行字符串，回车代表输入结束:")
// 	fmt.Scanln(&input)
// 	fmt.Println("输入的字符串：", input)
// }

package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
