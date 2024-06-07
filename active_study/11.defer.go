package main

import "fmt"

func main() {
	a()
	b()
}
func a() {
	i := 0
	// 在defer被声明的时候，就已经确定其值了
	defer fmt.Println(i) // 输出0
	i++
	return
}
func b() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
}
