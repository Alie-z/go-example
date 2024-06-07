package main

import "fmt"

func main() {
	i := 1
	switch i {
	case 0, 1:
		fmt.Println("i = ", i)

	case 3:
		fmt.Println("i = ", i)

	case 2:
		fmt.Println("i = ", i)
	}
}
