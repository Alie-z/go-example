package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	slice := intArr[1:3]
	fmt.Println("before change ...")
	fmt.Println("intArr =", intArr)
	fmt.Println("intArr[1]的地址为", &intArr[1])
	fmt.Println("slice[0]的地址为", &slice[0])
	fmt.Println("slice =", slice)
	fmt.Println("len(slice) =", len(slice))
	fmt.Println("cap(slice) =", cap(slice))

	fmt.Println("after change ...")
	slice[0] = 220
	fmt.Println("intArr =", intArr)
	fmt.Println("slice =", slice)
}
