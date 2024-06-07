package main

import "fmt"

func main() {
	a, b := multiResFun(1)
	fmt.Println(a, b)
	fmt.Println(multiResFun(-1))
}

func multiResFun(input int) (number int, err error) {
	number = input
	if input >= 0 {
		err = nil
		return
	}
	err = fmt.Errorf("number %d is negative", input)
	return
}
