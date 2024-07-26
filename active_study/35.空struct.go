// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		time.Sleep(1 * time.Second) // 执行一些操作...
// 		fmt.Printf("goroutine done\n")
// 		// close(done) // 不需要发送 struct{}{}，直接 close，发送完成信号
// 		done <- struct{}{}
// 	}()

// 	fmt.Printf("waiting...\n")
// 	<-done // 等待完成
// 	fmt.Printf("main exit\n")
// }

package main

import "fmt"

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

func main() {
	type A struct {
		noCopy noCopy
		a      string
	}

	type B struct {
		b string
	}

	a := A{a: "a"}
	b := B{b: "b"}

	c := a
	d := b
	fmt.Println(">>> c", c)
	fmt.Println(">>> d", d)
}
