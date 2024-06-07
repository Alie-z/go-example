package main

import "fmt"

func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}

	return false
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(Contain(slice, 3)) // 输出: true
	fmt.Println(Contain(slice, 6)) // 输出: false
}
