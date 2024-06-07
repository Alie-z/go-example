package main

import (
	"fmt"
	"sync"
	// 导入必要的包，例如 "sync" 用于互斥锁
	// 其他可能需要的包也可以用占位符代替
)

// 主函数
func main() {
	ConcurrentCount()
}

// 计数统计
var count int

// 读写锁
var rwLock sync.RWMutex

// GetCount 实现获取计数统计
func GetCount() int {
	// 读取数据的时候加读锁
	rwLock.RLock()
	defer rwLock.RUnlock()
	return count
}

// PlusOne 实现计数统计加一
func PlusOne() {
	// 写入数据的时候加写锁
	rwLock.Lock()
	defer rwLock.Unlock()
	count += 1
}

// ConcurrentCount 实现了多 goroutines 读写操作
func ConcurrentCount() {
	var wg sync.WaitGroup
	size := 100
	wg.Add(size)

	// 开启100个goroutines进行加一操作和读操作
	for i := 1; i <= size; i++ {
		go func() {
			defer wg.Done()
			PlusOne()
			GetCount()
		}()
	}
	wg.Wait()
	fmt.Printf("counter is: %d\n", count)
}
