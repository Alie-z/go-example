package main

import (
	"fmt"
	"sync"
	// 导入必要的包，例如 "sync" 用于互斥锁
	// 其他可能需要的包也可以用占位符代替
)

// 定义全局变量或常量（如果需要）

// 定义一个使用互斥锁的结构体（如果需要）

// 初始化函数或方法（如果需要）

// 主函数
func main() {
	NoLockSum()
	UseLockSum()
}

func NoLockSum() {
	var sum = 0
	var wg sync.WaitGroup
	size := 100
	wg.Add(size)
	for i := 1; i <= size; i++ {
		go func() {
			defer wg.Done()
			sum += 1
		}()
	}
	wg.Wait()
	fmt.Printf("sum is: %d\n", sum)
}
func UseLockSum() {
	var sum = 0
	var mutex sync.Mutex
	var wg sync.WaitGroup
	size := 100
	wg.Add(size)
	for i := 1; i <= size; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			sum += 1
		}()
	}
	wg.Wait()
	fmt.Printf("sum is: %d\n", sum)
}
